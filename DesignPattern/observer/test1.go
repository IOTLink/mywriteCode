package main

import (
	"strconv"

	"fmt"
)
func main() {
	m := make(map[int]string)
	for i:=0; i<10; i++ {
		m[i] = "hi" + strconv.Itoa(i)
	}

	for item, v := range m {
		fmt.Println(item," ", v)
	}

	return
}
