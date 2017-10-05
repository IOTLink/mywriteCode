package main

//rsa cipher

import (
	"crypto"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"encoding/base64"
	"fmt"
	"testing"

)

type Cipher interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
	Sign(src []byte, hash crypto.Hash) ([]byte, error)
	Verify(src []byte, sign []byte, hash crypto.Hash) error
}

func pkcs1Padding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize - 11

	var v [][]byte

	if srcSize <= blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]

			if len(src) < blockSize {
				v = append(v, src)
			}
		}
	}
	return v
}

func unPadding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize

	var v [][]byte

	if srcSize == blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]
		}
	}
	return v
}


type pkcsClient struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}



func (this *pkcsClient) Encrypt(plaintext []byte) ([]byte, error) {

	blocks := pkcs1Padding(plaintext, this.publicKey.N.BitLen()/8)

	buffer := bytes.Buffer{}
	for _, block := range blocks {
		ciphertextPart, err := rsa.EncryptPKCS1v15(rand.Reader, this.publicKey, block)
		if err != nil {
			return nil, err
		}
		buffer.Write(ciphertextPart)
	}

	return buffer.Bytes(), nil
}

func (this *pkcsClient) Decrypt(ciphertext []byte) ([]byte, error) {

	ciphertextBlocks := unPadding(ciphertext, this.privateKey.N.BitLen()/8)

	buffer := bytes.Buffer{}
	for _, ciphertextBlock := range ciphertextBlocks {
		plaintextBlock, err := rsa.DecryptPKCS1v15(rand.Reader, this.privateKey, ciphertextBlock)
		if err != nil {
			return nil, err
		}
		buffer.Write(plaintextBlock)
	}

	return buffer.Bytes(), nil
}

func (this *pkcsClient) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, hashed)
}

func (this *pkcsClient) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(this.publicKey, hash, hashed, sign)
}

type Type int64

const (
	PKCS1 Type = iota
	PKCS8
)


//默认客户端，pkcs8私钥格式，pem编码
func NewDefault(privateKey, publicKey string) (Cipher, error) {
	blockPri, _ := pem.Decode([]byte(privateKey))
	if blockPri == nil {
		return nil, errors.New("private key error")
	}

	blockPub, _ := pem.Decode([]byte(publicKey))
	if blockPub == nil {
		return nil, errors.New("public key error")
	}

	return New(blockPri.Bytes, blockPub.Bytes, PKCS8)
}

func New(privateKey, publicKey []byte, privateKeyType Type) (Cipher, error) {

	priKey, err := genPriKey(privateKey, privateKeyType)
	if err != nil {
		return nil, err
	}
	pubKey, err := genPubKey(publicKey)
	if err != nil {
		return nil, err
	}
	return &pkcsClient{privateKey: priKey, publicKey: pubKey}, nil
}

func genPubKey(publicKey []byte) (*rsa.PublicKey, error) {
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func genPriKey(privateKey []byte, privateKeyType Type) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, errors.New("unsupport private key type")
		}
	}
	return priKey, nil
}



var cipher Cipher

func init() {
	client, err := NewDefault(`
-----BEGIN PRIVATE KEY-----
私钥信息
-----END PRIVATE KEY-----`,
`-----BEGIN PUBLIC KEY-----
公钥信息
-----END PUBLIC KEY-----
`)

	if err != nil {
		fmt.Println(err)
	}

	cipher = client
}

func Test_DefaultClient(t *testing.T) {

	cp, err := cipher.Encrypt([]byte("测试加密解密"))
	if err != nil {
		//t.Error(err)
	}
	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		//t.Error(err)
	}
	pp, err := cipher.Decrypt(ppBy)

	fmt.Println(string(pp))
}

func Test_Sign_DefaultClient(t *testing.T) {

	src := "测试签名验签"

	signBytes, err := cipher.Sign([]byte(src), crypto.SHA256)
	if err != nil {
	//	t.Error(err)
	}
	sign := hex.EncodeToString(signBytes)
	fmt.Println(sign)

	signB, err := hex.DecodeString(sign)

	errV := cipher.Verify([]byte(src), signB, crypto.SHA256)
	if errV != nil {
		//t.Error(errV)
	}
	fmt.Println("verify success")
}


func main() {
	t := &testing.T{}
	Test_DefaultClient(t)
	Test_Sign_DefaultClient(t)

}

//https://my.oschina.net/u/1023800/blog/526936
//cipher 封装思想
//不同i标准


