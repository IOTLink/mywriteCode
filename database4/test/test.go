package main

import "fmt"

type DB struct {
	*SubDB
}

type SubDB struct {
	string
}

func (sub *SubDB) GetString() string{
	return sub.string
}

func (sub *SubDB) SetString(s string) {
	sub.string = s
}

func test1() {
	var sub SubDB
	var s DB
	s.SubDB = &sub
	s.SetString("hello world")
	fmt.Println(s.GetString())
}

func test2() {
	var s DB   //调用的是子类的方法　，因为子类没有赋值虚类　　会运行崩溃
	s.SetString("hello world")
	fmt.Println(s.GetString())
}

func main() {
	//test1()
	test2()
}