package main

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	pwd := []byte("adminw")
	b, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
	fmt.Println(string(b))


	pwd = []byte("adminw")
	b, err = bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
	fmt.Println(string(b))

	err = bcrypt.CompareHashAndPassword(b,pwd)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("same key")
}