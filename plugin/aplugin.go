
package main

func Add(x, y int) int {
	return x+y
}

func Subtract(x, y int) int {
	return x-y
}

//go build -buildmode=plugin -o aplugin.so aplugin.go