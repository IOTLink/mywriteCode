package main
//import "crypto/sha1"
//import "fmt"

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func sha1test(){
	s := "sha1 this string"

	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

func sha256test() {
	s := "this a sha256 string"
	h := sha256.New()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	// a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", hash)
}


func sha512test() {
	s := "this a sha512 string"
	h := sha512.New()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	// a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", hash)
}

func md5test() {
	s := "this a md5 string"
	h := md5.New()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%s\n", hash)
	fmt.Printf("%x\n", hash)
	fmt.Printf("%x\n", hex.EncodeToString(hash))
}

func main() {
	sha1test()
	sha256test()
	sha512test()
	md5test()
}
