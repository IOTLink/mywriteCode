package main

import (
	"fmt"
	//"sync"
	"time"
)

func main() {
	test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
}

func test1() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()

	msg :=<-messages
	fmt.Println(msg)
	return
}

func test2() {
	exitChan := make(chan int,1)
	msgPing := make(chan string)
	msgPong := make(chan string)

	go func() {
		count := 0
		fmt.Println("go groutunue 1")
		for {
			msgPing <- "ping"

			select {
			case msg := <-msgPong:
				fmt.Println(msg)
			}

			if count > 3 {
				break
			} else {
				count++
			}
			time.Sleep(time.Millisecond*500)
		}
		exitChan <- 1
	}()

	go func() {
		fmt.Println("go groutunue 2")
		for{
			count := 0
			select {
			case msg := <-msgPing:
				fmt.Println(msg)
			//default:
			//	fmt.Println("go 2 continue")
			}
			msgPong <- "pong"
			if count > 3 {
				break
			} else {
				count++
			}

			time.Sleep(time.Millisecond*500)
		}

		exitChan <- 2
	}()

	time.Sleep(time.Second*1)
	return
	for i:=0; i<2; i++ {
		exit := <-exitChan
		fmt.Println(exit)
		/*select {
			case exit := <-exitChan:
			fmt.Println(exit)
		}*/
	}
	//close(msgPing)
}


func test3() {
	exitChan := make(chan int,1)
	msgPing := make(chan string)
	msgPong := make(chan string)
	go func() {

		for {
			msgPing <- "ping"
			msgPong <- "pong"
			time.Sleep(time.Millisecond*500)
		}
	}()

	go func() {
		fmt.Println("go groutunue 1")
		time.Sleep(time.Second*1)
		for {
			select {
			case msg := <-msgPong:
				fmt.Println(msg)
			}

		}
		exitChan <- 1
	}()

	go func() {
		fmt.Println("go groutunue 2")
		time.Sleep(time.Second*1)

		for{
			select {
			case msg := <-msgPing:
				fmt.Println(msg)
			}
		}

		exitChan <- 2
	}()

	time.Sleep(time.Second*2)
	close(msgPing)
	close(msgPong)
}

func test4() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func test5(){
	fmt.Println("hello world")
	select {
	case <-time.After(time.Millisecond*500):
		fmt.Println("time out")
	}
	/*
	select {

	}
	*/
}

func test6() {
	fmt.Println("start parmages!")
	go func(){
		for {
			time.Sleep(time.Millisecond* 500)
			fmt.Println("sleep time out")
		}
	}()
	select {

	}
}


func test8() {
	fmt.Println("start parmages!")
	exitChan := make(chan int,1)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			fmt.Println("sleep time out")

			select {
			case <-exitChan:
				break
			default:
				fmt.Println("connection")
			}
		}
	}()
	select {
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
		exitChan <-1
	}

}

func test7() {
	timeout := make(chan bool, 1)
	go func() {
		//fmt.Println("test go time out")
		time.Sleep(time.Millisecond * 500)
		timeout <- true
	}()

	select {
		case <-timeout:
			fmt.Println("test out")
	}
}



