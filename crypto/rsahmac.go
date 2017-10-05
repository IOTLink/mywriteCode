package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
	"crypto"
)

func main() {
	//rsa 密钥文件产生
	GenRsaKey(1024)
}
//RSA公钥私钥产生
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}


	//test hamc demo
	h := sha1.New()
	io.WriteString(h, "hello world......")
	hash0 := h.Sum(nil)
	fmt.Printf("%x\n", hash0)

	//hmac ,use sha1
	key := derStream//derPkix
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("hello world......"))
	sign1 := mac.Sum(nil)
	fmt.Printf("%x\n", sign1) //使用sha1对信息做摘要，在使用私钥加密摘要信息，作为签名

	//??
	sign ,err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash0)
	fmt.Printf("%x\n", sign)
	//sha1hash, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, hash)
	//fmt.Printf("%x\n", sha1hash)



	return nil
}

