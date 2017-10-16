package main

import (
	"crypto/aes"
	"fmt"
	"bytes"
)



// PKCS7Padding pads as prescribed by the PKCS7 standard
func PKCS7Padding(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize  // 16 - 20%16 = 16 - 4 = 12
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)  // 12 字面值。重复12次
	return append(src, padtext...) //最佳到源字符串后后面
}

// PKCS7UnPadding unpads as prescribed by the PKCS7 standard
func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > aes.BlockSize || unpadding == 0 {
		return nil, fmt.Errorf("invalid padding")
	}

	pad := src[len(src)-unpadding:]
	for i := 0; i < unpadding; i++ {
		if pad[i] != byte(unpadding) {
			return nil, fmt.Errorf("invalid padding")
		}
	}

	return src[:(length - unpadding)], nil
}


func main() {
	src := []byte("12345678901234567890") //= 20 byte
	srctmp := PKCS7Padding(src)
	fmt.Println(srctmp)

	ss, err := PKCS7UnPadding(srctmp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ss)
}

/*
加密中常用的填充字符byte


 */