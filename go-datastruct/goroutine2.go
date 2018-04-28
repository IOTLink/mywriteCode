package main

import (
	"sync"
	"fmt"
	//"time"
	//"time"
)

const (
	go_size int = 10
	count int = 10000
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(go_size)

	wgWorker := sync.WaitGroup{}
	wgWorker.Add(1)

	msg := make(chan string,go_size)
	done := make(chan bool,1)

	go func() {
		quit := false
		for {
			
			select {
				case msg <- "ping":
					//fmt.Println("send ping")
				case close := <-done: {
						fmt.Println("close main", close)
					    quit = true
						break
					}
				default :
			}

			if quit {
				break
			}
		}
		wgWorker.Done()
	}()

	for i:=0; i<go_size; i++{
		go func(id int, oper int) {
			count := 0
			for i:=0; i<oper; i++{
				select {
				case message := <-msg:
					count += 1

					if count == oper {
						fmt.Println(id, " recv ", message, " ", count)
					}
				}
			}
			wg.Done()
		}(i, count)
	}

	wg.Wait()
	//close(msg)
	//time.Sleep(time.Second*500)
	/*
	for Msg := range msg {
		fmt.Println(Msg)
		done <-true
	}
	*/
	//<-msg
	done <-true
	//time.Sleep(time.Second*2)
	wgWorker.Wait()
}
