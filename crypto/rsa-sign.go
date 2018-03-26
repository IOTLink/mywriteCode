package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"os"
	"fmt"
	"hash"
	"crypto/sha256"
)

func GetPrivateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil
	}
	return privateKey
}

func GetPublicKeyFromPrivateKey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	publicKey := &privateKey.PublicKey
	return publicKey
}

func rsaSign() error{

	//rsa 在pkcs1中规定
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	var pubkey rsa.PublicKey
	pubkey = privateKey.PublicKey


	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privateKey)

	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubkey)


	var h hash.Hash
	h = sha256.New()
	data := []byte("This is a message to be signed and verified by AES!")
	h.Write(data)

	//io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
	//hashed := h.Sum(nil)
	hashed := h.Sum(nil)



	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Signature : %x\n", signature)


	//verify
	var h1 hash.Hash
	h1 = sha256.New()
	data1 := []byte("This is a message to be signed and verified by AES!")
	h1.Write(data1)

	hashed1 := h1.Sum(nil)
	err =  rsa.VerifyPKCS1v15(&pubkey, crypto.SHA256, hashed1[:], signature)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("verify success")

	return nil
}


func main() {
	rsaSign()
}