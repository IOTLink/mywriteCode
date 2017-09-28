package main


import "encoding/pem"
import "encoding/asn1"
import (
	"crypto/dsa"
	"fmt"
	"crypto/rand"
	"crypto/x509"
	"crypto/sha1"
)

var signingPubKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIICIDANBgkqhkiG9w0BAQEFAAOCAg0AMIICCAKCAgEApSmU3y4DzPhjnpOrdpPs
cIosWJ4zSV8h02b0abLW6nk7cnb5jSwBZKLrryAlF4vs+cF1mtMYjX0QKtEYq2V6
WVDnoXj3BeLYVbhsHuvxYmwXmAkNsSnhMfSCxsck9y6zuNeH0ovzBD90nISIJw+c
VAnUt0dzc7YKjBqThHRAvi8HoGZlzB7Ryb8ePSW+Mfr4jcH3Mio5T0OH3HTavN6Y
zpnohzQo0blwtwEXZOwrNPjQNrSigdPDrtvM32+hLTIJ75Z2NbIRLBjNlwznu7dQ
Asb/AiPTHXihxCRDm+dH70dps5JfT5Zg9LKsPhANk6fNK3e4wdN89ybQsBaswp9h
xzORVD3UiG4LuqP4LMCadjoEazShEiiveeRBgyiFlIldybuPwSq/gUuFveV5Jnqt
txNG6DnJBlIeYhVlA25XDMjxnJ3w6mi/pZyn9ZR9+hFic7Nm1ra7hRUoigfD/lS3
3AsDoRLy0xZqCWGRUbkhlo9VjDxo5znjv870Td1/+fp9QzSaESPfFAUBFcykDXIU
f1nVeKAkmhkEC9/jGF+VpUsuRV3pjjrLMcuI3+IimfWhWK1C56JJakfT3WB6nwY3
A92g4fyVGaWFKfj83tTNL2rzMkfraExPEP+VGesr8b/QMdBlZRR4WEYG3ObD2v/7
jgOS2Ol4gq8/QdNejP5J4wsCAQM=
-----END PUBLIC KEY-----
`)

func test1(){
	block, _ := pem.Decode(signingPubKey)
	if block == nil {
		fmt.Errorf("expected block to be non-nil", block)
		return
	}

	var pubkey dsa.PublicKey

	_,err := asn1.Unmarshal(block.Bytes, &pubkey)
	if err != nil {
		fmt.Errorf("could not unmarshall data: `%s`", err)
	}

	fmt.Printf("public key param P: %d\n", pubkey.Parameters.P)
	fmt.Printf("public key param Q: %d\n", pubkey.Parameters.Q)
	fmt.Printf("public key param G: %d\n", pubkey.Parameters.G)
	fmt.Printf("public key Y: %d\n", pubkey.Y)

	fmt.Printf("done")
}

func generatePrivKey() *dsa.PrivateKey{
	params := dsa.Parameters{}
	err := dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160)
	if err != nil {
		fmt.Errorf("failed to generate dsa parameters: %v", err)
		return nil
	}

	// Create the DSA private/public keys
	priv := new(dsa.PrivateKey)
	priv.Parameters = params
	err = dsa.GenerateKey(priv, rand.Reader)
	if err != nil {
		fmt.Errorf("failed to generate dsa keys: %v", err)
		return nil
	}
	return priv
}

func test2(){
	var message = "Hello brave new world!"
	var hash = sha1.Sum([]byte(message))
	var err error

	priv := generatePrivKey()

	// Sign a message
	r, s, err := dsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		fmt.Errorf("failed to sign message: %v", err)
	}

	if !dsa.Verify(&priv.PublicKey, hash[:], r, s) {
		fmt.Errorf("failed to verify message: %v", err)
	}


	//private to  pem block

}
func main() {
	test2()
}
/*
https://hsulei.com/2016/10/12/%E6%95%B0%E5%AD%97%E8%AF%81%E4%B9%A6%E5%92%8Cgolang%E7%9A%84%E7%A0%94%E7%A9%B6/

同在x509包下提供了:

1
2
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
把RSA和ECDSA私钥转换成byte数组的方法，但是没有找到把DSA私钥转换成byte数组的方法。

 */