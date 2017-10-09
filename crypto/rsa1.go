package main

import (
	"crypto/x509"
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"errors"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha256"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDJZmXxiygEXQQoKMps8T4bYyw3Nt4V6EnF8//aieeTRoDweP3f
rwRLri9fnOtF9U64tWEvZi9zqqcW4q70cW0VQw3ZcqSLGYhtmCWCSwpgLILZ4UO2
CNmnbT83imqXCgf7XkreE6FvNez/0k7hHxp9TtSoD8sKEVO8SCnFvLigOQIDAQAB
AoGBAKrAknXz0X3BjyCtVNKAxnNZ5cb6UdipFMMsWCPk/92xzOgU7MmVDUNM7wVV
eXPaw19/CGKzdE/Qw9F3Z3YAJMZ8+T8+eGX4ERGE53rCgtONSS7V67UhWT0GCHf9
zAU4qcLoPNb383yNKHc5b6ZZwSOVfuegd9bDuMxZ+XDjEp9xAkEA5wdQLe8q7YE+
w9wmrOwq+I5HvnLQTNuBgxHx7Gx0Lmv3dBTB+xZA9mK+VC3I3tScl5d5Bmqk+xEJ
uHwqxX5VtwJBAN8rPwS5Vf82ck7RVH6rpHWiNc9R7LIaksw9rtsamGogiqj0ac7j
qCw9cPBp9vraPr9ya6c0586qRO589R53OY8CQQCNc17WIP+Jl36UcPCFI7xTntmy
c52x0RwE4jHbbbPc0GIcArGaSE/SCzc5VycLt+WAs094bEdDKXVoLS4K6YUVAkB4
JnxjOrVGFjYsoR7wo9CDVLXLzLu2l//43izJ4eO1H5gHpq9gp8jfKKUpiqNiIyYt
RJkZCK7U4W8DR1tg76eZAkEA2s464nE3bpY9KIeNo9mohFf/2vxcAAoHKDwyjCSA
E+EW6at9dAGsQbiwEh0uneNLmuippM+1xPMggTxQjDxhxQ==
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJZmXxiygEXQQoKMps8T4bYyw3
Nt4V6EnF8//aieeTRoDweP3frwRLri9fnOtF9U64tWEvZi9zqqcW4q70cW0VQw3Z
cqSLGYhtmCWCSwpgLILZ4UO2CNmnbT83imqXCgf7XkreE6FvNez/0k7hHxp9TtSo
D8sKEVO8SCnFvLigOQIDAQAB
-----END PUBLIC KEY-----
`)


var publicKey1 = []byte(`
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJZmXxiygEXQQoKMps8T4bYyw3
Nt4V6EnF8//aieeTRoDweP3frwRLri9fnOtF9U64tWEvZi9zqqcW4q70cW0VQw3Z
cqSLGYhtmCWCSwpgLILZ4UO2CNmnbT83imqXCgf7XkreE6FvNez/0k7hHxp9TtSo
D8sKEVO8SCnFvLigOQIDAQAB
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}



func main() {
	data, err := RsaEncrypt([]byte("polaris@studygolang.com"))
	if err != nil {
		panic(err)
	}
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))


	//测试私钥做签名信息 验证 测试
	message := []byte("hello world..................")
	mac := hmac.New(sha256.New, publicKey1)
	mac.Write(message)
	sig := mac.Sum(nil)
	fmt.Printf("sig %x\n", sig)

	hash1,err := RsaDecrypt(sig)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Printf("hash1 %x\n", hash1)

	h := sha256.New()
	h.Write([]byte(message))
	hash2 := h.Sum(nil)
	fmt.Printf("hash2 %x\n", hash2)

	if len(hash1) != len(hash2) {
		fmt.Println("length is not same")
	}
}


//http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Brsa/
/*

1）创建私钥：
openssl genrsa -out private.pem 1024 //密钥长度，1024觉得不够安全的话可以用2048，但是代价也相应增大
2）创建公钥：
openssl rsa -in private.pem -pubout -out public.pem

 */