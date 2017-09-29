package main

import (
	"fmt"
	"reflect"
)

func test1() {
	var a interface{}
	var i int = 5
	s := "Hello world"
	//var s string
	//s = "hello world"
	// a可以存储任意类型的数值
	a = i
	fmt.Println(a)
	v, ok := a.(int)
	if ok {
		fmt.Println(ok, v)
	}

	a = s
	fmt.Println(a)

	v1, ok1 := a.(string)
	if ok1 {
		fmt.Println(ok1, v1)
	}


	dty := reflect.TypeOf(a)
	fmt.Println(dty)
}

func testX() {
	var a interface{}
	a = 555
	fmt.Println(a)
	a = "hello world"
	value, ok := a.(string)
	if !ok {
		fmt.Println("It's not ok for type string")
		return
	}
	fmt.Println("The value is ", value)

}

//_______________________________________________________

//定义了一个接口
type I interface {
	Get() int
	Put(int)
}

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

func f1(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

//interface{}空接口，能接受任何类型。.(I)是类型断言.
func f2(p interface{}) {
	if t, ok := p.(S); ok {
		fmt.Println("S:", t)
	} else if t, ok := p.(I); ok {
		fmt.Println("I:", t.Get())
	}
}

func f3(p interface{}) {
	switch t := p.(type) {
	case S:
		fmt.Println("S:", t.Get())
	case R:
		fmt.Println("R:", t.Get())
	case I:
		fmt.Println("I:", t.Get())
	default:
		fmt.Println("unknow type")
	}
}

func test2(){
	s := S{101}

	f1(&s)
	f2(&s)

	r := R{1111}
	f3(&r)

	var i1 I
	i1 = &s

	fmt.Println("i1:", i1.Get())
	i1.Put(8888)
	fmt.Println("i1:", i1.Get())

	f2(i1)
	f3(i1)

}



func main() {
	test1()
	test2()
	testX()
}
