package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
)
func main() {
	//sha1
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))


	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	// mac := hmac.New(md5.New, key)
	mac.Write([]byte("aaaaaa"))
	hash := mac.Sum(nil)
	fmt.Printf("%x\n", hash)
	fmt.Printf("%s\n", string(hash))


	//接受者验证
	ret := hmac.Equal(hash, hash)
	fmt.Println(ret)

	/*
	mac.Write(hash)
	hash = mac.Sum(nil)
	fmt.Printf("%x\n", hash)
	*/

}

//http://www.nljb.net/default/Golang%E4%B9%8BHMAC%E7%9A%84SHA1%E5%8A%A0%E5%AF%86/
