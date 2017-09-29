package main

import (
	"fmt"
)


type I interface {
	Get() int
	Set(int)
}
type SS struct {
	Age int
}
func (s SS) Get() int {
	return s.Age
}
func (s SS) Set(age int) {
	s.Age = age
}
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func test1() {
	ss := SS{Age:10000}
	f(&ss) //ponter //go 会把指针进行隐式转换得到 value，但反过来则不行。
	f(ss)  //value

}








func main(){

}


//http://sanyuesha.com/2017/07/22/how-to-understand-go-interface/
