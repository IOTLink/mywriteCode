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
	for {
		select {
		case req := <-queue:
			go handle(req)
		case <-time.After(time.Millisecond * 500):
			fmt.Println("time out")
		}
	}
}

func main(){
	queue := make(chan *Request,10)

	for i:=0; i<10; i++{
		r := &Request{key:i,value:string("he!")}
		queue <- r
		fmt.Println("instert queue",i)
	}
	Serve(queue)

}