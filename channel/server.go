package main

import (
	"fmt"
	//"sync"
	"time"

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
	channel_close_flag := false
	for {
		if channel_close_flag {
			return
		}
		select {
		case req ,ok := <-queue:
			if !ok {
				fmt.Println("exit 0")
				channel_close_flag = true
				break
			}
			go handle(req)
		case <-time.After(time.Millisecond * 500):
			fmt.Println("time out")
		}
	}
	*/

	var sem = make(chan int, 1)
	for req := range queue{
		//fmt.Println("queue is open")
		sem<-1
		go func(r *Request) {
			process(r)
			<-sem
		}(req)

	}

	/*
	for req := range queue{
		process(req)
	}
	*/
}

func test1() {
	queue := make(chan *Request,100)

	for i:=0; i<100; i++{
		r := &Request{key:i,value:string("he!")}
		queue <- r
		fmt.Println("instert queue",i)
	}
	close(queue)
	Serve(queue)
}

func main(){
	test1()
	time.Sleep(time.Second*3)
}

