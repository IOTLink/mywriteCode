package main

import (
	"fmt"
	//"unsafe"
)


type S struct{}

func (s *S) addr() { fmt.Printf("%p\n", s) }


func main(){

	var a, b S
	a.addr() // 0x1beeb0
	b.addr() // 0x1beeb0



}
