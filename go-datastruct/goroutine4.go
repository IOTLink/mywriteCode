package main

import (
	//"sync"
	"fmt"
	//"time"
	//"time"
	"time"
)

func main() {
	msg1:= make(chan string,1)
	//msg1:= make(chan string)
	msg2:= make(chan string, 1)
	done := make(chan bool)

	go func() {
		quit := false
		for {
			select {
				case msg1 <- "ping":
					fmt.Println("msg1 send ping")
					//time.Sleep(time.Millisecond*200)
				case msg2 <- "ping":
					fmt.Println("msg2 send pong")

				case msg :=<- msg1:
					fmt.Println("msg1 recv ",msg)
					//time.Sleep(time.Millisecond*200)
				case msg :=<- msg2:
					fmt.Println("msg2 recv ",msg)
					//time.Sleep(time.Millisecond*200)

				case close := <-done: {
						fmt.Println("close main", close)
					    quit = true
						break
					}
				default :
					break
				//	fmt.Println("default 1")
			}

			if quit {
				break
			}
		}
	}()

	time.Sleep(time.Second*200)
	done <-true
}

/*
无缓冲： 不仅仅是向 c1 通道放 1，而是一直要等有别的携程 <-c1 接手了这个参数，那么c1<-1才会继续下去，要不然就一直阻塞着。
有缓冲： c2<-1 则不会阻塞，因为缓冲大小是1(其实是缓冲大小为0)，只有当放第二个值的时候，第一个还没被人拿走，这时候才会阻塞。
*/
