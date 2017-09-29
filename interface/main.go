package main

import (
	"fmt"
)

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

/*
func (t *T) M1() {
	t.Name = "name1"
}
func (t T) M2() {
	t.Name = "name2"
}
*/
func (t *T) M2() {
	t.Name = "name2"
}

func test1() {
	t1 := T{"t1"}

	fmt.Println("M1调用前：", t1.Name)
	t1.M1()
	fmt.Println("M1调用后：", t1.Name)

	fmt.Println("M2调用前：", t1.Name)
	t1.M2()
	fmt.Println("M2调用后：", t1.Name)

	//________________________

	t2 := &T{"t2"}

	fmt.Println("M1调用前：", t2.Name)
	t2.M1()
	fmt.Println("M1调用后：", t2.Name)

	fmt.Println("M2调用前：", t2.Name)
	t2.M2()
	fmt.Println("M2调用后：", t2.Name)

}


type Intf interface {
	M1()
	M2()
}

func fn(t Intf) {
	t.M1()
	t.M2()
}


func test2() {

	var t1 T = T{"t1"}
	fmt.Println("M1调用前：", t1.Name)
	t1.M1()
	fmt.Println("M1调用后：", t1.Name)

	fmt.Println("M2调用前：", t1.Name)
	t1.M2()
	fmt.Println("M2调用后：", t1.Name)

	//var t2 Intf = t1
	var t2 Intf = &t1
	t2.M1()
	t2.M2()

	fn(&t1)
}



func main() {
	//test1()
	test2()
}
//https://studygolang.com/articles/4059

