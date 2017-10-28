package main

import (
	"fmt"
	"unsafe"
)


func main() {
	var s string
	var c complex128
	fmt.Println(unsafe.Sizeof(s))    // prints 8
	fmt.Println(unsafe.Sizeof(c))    // prints 16


	var a [3]uint32
	fmt.Println(unsafe.Sizeof(a)) // prints 12


	type S struct {
		a uint16
		b uint32
	}
	var s1 S
	fmt.Println(unsafe.Sizeof(s1)) // prints 8, not 6

	var s2 struct{}
	fmt.Println(unsafe.Sizeof(s2)) // prints 0


	type S3 struct {
		A struct{}
		B struct{}
	}
	var s3 S3
	fmt.Println(unsafe.Sizeof(s3)) // prints 0


	var x1 [1000000000]struct{}
	fmt.Println("[1000000000]struct{}",unsafe.Sizeof(x1)) // prints 0


	var x2 = make([]struct{}, 1000000000)
	fmt.Println("make([]struct{}, 1000000000)",unsafe.Sizeof(x2)) // prints 12 in the playground


	var x = make([]struct{}, 100)
	var y = x[:50]
	fmt.Println(len(y), cap(y)) // prints 50 100


	//var a5 struct{}
	//var b5 = &a5

	var a6, b6 struct{}
	fmt.Println(&a6 == &b6) // true


	a8 := make([]struct{}, 10)
	b8 := make([]struct{}, 20)
	fmt.Println(&a8 == &b8)       // false, a and b are different slices
	fmt.Println(&a8[0] == &b8[0]) // true, their backing arrays are the same

	a9 := struct{}{} // not the zero value, a real new struct{} instance
	b9 := struct{}{}
	fmt.Println(a9 == b9) // true





}
