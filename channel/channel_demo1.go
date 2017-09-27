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
	go func() {
		for {
			select {
			case req := <-queue:
				go handle(req)
			case <-time.After(time.Millisecond * 500):
				fmt.Println("time out")
			}
		}
		//req := <-queue
		//go handle(req)  // Don't wait for handle to finigo.
	}()
}

func main(){
	queue := make(chan *Request,10)
	Serve(queue)
	for i:=0; i<10; i++{
		r := &Request{key:i,value:string("he!")}
		queue <- r
		fmt.Println("instert queue",i)
	}

	time.Sleep(time.Second *100)
	select {
	}
}