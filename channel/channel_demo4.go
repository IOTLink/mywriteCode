package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	test2()

}

func test2() {
	var wg sync.WaitGroup
	msgPing := make(chan string)
	msgPong := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		count := 0
		fmt.Println("go groutunue 1")

		for {
			msgPing <- "ping"

			select {
			case msg ,ok := <-msgPong:
				if !ok {
					return
				}
				fmt.Println(msg,count)
			}

			if count > 3 {
				break
			} else {
				count++
			}
			time.Sleep(time.Millisecond*500)
		}
		close(msgPing)

	}()

	go func() {
		fmt.Println("go groutunue 2")
		defer wg.Done()
		count := 0
		for{

			select {
			case msg, ok := <-msgPing:
				if !ok {
					return
				}
				fmt.Println(msg,count)
			}
			msgPong <- "pong"
			if count > 3 {
				break
			} else {
				count++
			}
			time.Sleep(time.Millisecond*500)
		}
		close(msgPong)
	}()
	wg.Wait()

}

