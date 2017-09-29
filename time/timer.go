package main

import (
	"time"
	"fmt"
)

func test1() {
	for {
		t1 := time.Now()
		fmt.Println(t1)
		if t1.IsZero() {
			fmt.Println(t1)
		}
	}
}

func main() {
	test1()
}
