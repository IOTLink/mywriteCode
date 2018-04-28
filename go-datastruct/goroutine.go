package main

import (
	"sync"
	"fmt"
	//"time"
)

const (
	go_size int = 3
	count int = 10
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(go_size)

	msg := make(chan string,go_size)
	done := make(chan bool,1)
	go func() {
		for {

			select {
				case msg <- "ping":
					{

					}
				case close := <-done:
					fmt.Println("close main", close)
					return
				default:
					//fmt.Println("default")
			}
		}
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
	//time.Sleep(time.Second*500)
	done <-true
}
