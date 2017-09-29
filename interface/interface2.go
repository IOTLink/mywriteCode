package main

import (
	"fmt"
)

type I interface {
	Get() int
	Set(int)
}
//2
type S struct {
	Age int
}



func(s S) Get()int {
	return s.Age
}
func(s *S) Set(age int) {
	s.Age = age
}

type R struct {
	Age int
}

func(s R) Get()int {
	return s.Age
}
func(s *R) Set(age int) {
	s.Age = age
}



//3
func f(i I){
	i.Set(10)
	fmt.Println(i.Get())
}


func printAll(vals []interface{}) { //1
	for _, val := range vals {
		fmt.Println(val)
	}
}


func main() {
	s := S{}
	f(&s)  //4

	s1 := S{Age:1000}
	var i I //声明 i
	i = &s1 //赋值 s 到 i
	fmt.Println(i.Get())


	if t, ok := i.(*S); ok {
		fmt.Println("s implements I", t)
	}


	switch t := i.(type) {
	case *S:
		fmt.Println("i store *S", t)
	case *R:
		fmt.Println("i store *R", t)
	}

	names := []string{"stanley", "david", "oscar"}
	var interfaceSlice []interface{} = make([]interface{}, len(names))
	for i, d := range names[:] {
		interfaceSlice[i] = d
	}
	printAll(interfaceSlice)

}


