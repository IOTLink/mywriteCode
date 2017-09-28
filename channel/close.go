package main

import (
	"fmt"
	"sync"
)
type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func test1(){
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true

}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}


func SafeSend(ch chan T, value T) (closed bool) {

	defer func() {
		if recover() != nil {
			// the return result can be altered
			// in a defer function call
			closed = true
		}
	}()

	ch <- value // panic if ch is closed
	return false // <=> closed = false; return
}

func test2() {
	c := make(chan T)
	close(c)
	ret := SafeSend(c,10)
	//close(c)
	fmt.Println(ret)
}



func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true
}

func test3() {
	c := make(chan T)
	close(c)
	ret := SafeSend(c,10)
	SafeClose(c)
	fmt.Println(ret)
}



type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func(){
		close(mc.C)
	})
}

func test4(){
	c := NewMyChannel()
	c.SafeClose()
	c.SafeClose()

}