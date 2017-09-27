package main

import (
	"fmt"
	//"sync"
	//"time"

)

type Request struct{
	key int
	value string
}

func process(r *Request) {
	//time.Sleep(time.Millisecond* 500)
	fmt.Println("hello world",r.key, r.value)
}

func handle(r *Request) {
	process(r)  // May take a long time.
}

func Serve(queue chan *Request) {
	/*
	for {
		select {
		case req := <-queue:
			go handle(req)
		case <-time.After(time.Millisecond * 500):
			fmt.Println("time out")
		}
	}
	*/
	for req := range queue{
		go func(r *Request) {
			process(r)
		}(req)
	}
}

func test1() {
	queue := make(chan *Request,10)

	for i:=0; i<10; i++{
		r := &Request{key:i,value:string("he!")}
		queue <- r
		fmt.Println("instert queue",i)
	}
	Serve(queue)
}

func main(){
	//test1()
	test2()

}

func test2() {
	queue := make(chan int,100)
	for i:=0; i<10; i++{
		queue <- i
		fmt.Println("instert queue",i)
	}

	for req := range queue{
		fmt.Println(req)
	}
}
//http://www.jianshu.com/p/fe5dd2efed5d
//https://gobyexample.com/range-over-channels
//go range channel