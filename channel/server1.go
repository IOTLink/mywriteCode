package main

import (
	"fmt"
	//"sync"
	//"time"

)

const (
	MaxOutstanding int = 10
)
type Request struct{
	key int
	value string
}

func process(r *Request) {
	//time.Sleep(time.Millisecond* 500)
	fmt.Println("hello world",r.key, r.value)
}

func handle(queue chan *Request) {
	for req := range queue{
		process(req)  // May take a long time.
	}

}

func Serve(queue chan *Request,quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handle(queue)
	}
	<-quit
}

func test1() {
	quit := make(chan bool,MaxOutstanding)
	queue := make(chan *Request,100)

	for i:=0; i<100; i++{
		r := &Request{key:i,value:string("he!")}
		queue <- r
		fmt.Println("instert queue",i)
	}
	//quit<-true
	close(quit)
	close(queue)
	Serve(queue,quit)
}

func main(){
	//test1()
	//time.Sleep(time.Second*3)

	test2()
}

func test2() {
	x := make(chan int)
	go func() {
		x <- 1
	}()
	<-x
}