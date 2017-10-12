package main

import (
	"time"
	"fmt"
)
func main() {

	dst := make([]byte,100)
	src := "hello world"

	start := time.Now()
	for i:=0; i<3000000; i++{
		copy(dst,src)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("copy time:",elapsed)


	start = time.Now()
	for i:=0; i<3000000; i++{
		dst = append(dst, src...)
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("append time:",elapsed)

	var a struct {}
	fmt.Println("struct 0:",a)
}