package main

import (
	"fmt"
	"reflect"
)

type Ele interface{}

type List []Ele

func main() {
	list := make(List, 4)
	list[0] = 1
	list[1] = 'c'
	list[2] = "string"
	list[3] = [2]int{5, 6}
	for index, val := range list {
		switch typeval := val.(type) {
		case int:
			fmt.Printf("list[%d] is an int(%d)\n", index, typeval)
		case string:
			fmt.Printf("list[%d] is a string(%s)\n", index, typeval)
		case rune:
			fmt.Printf("list[%d] is a rune(%c)\n", index, typeval)
		default:
			fmt.Printf("list[%d] is a different type(%s)\n", index, reflect.TypeOf(typeval))
		}
	}
}

/*
list[0] is an int(1)
list[1] is a rune(c)
list[2] is a string(string)
list[3] is a different type([2]int)

*/