package main

import (
	"net"
	"bufio"
	"strconv"
	"fmt"
)

const (
	PORT = 7051
	MAX_CONNECT_LEN = 100000
)

func deal(client net.Conn) {
	defer client.Close()
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}

func main(){
	var i uint64
	var CONNECT_CHANNEL = make(chan net.Conn, MAX_CONNECT_LEN)

	listener, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
	if err != nil {
		fmt.Printf("couldn't start listening: %s",err.Error())// + err.String())
		return
	}
	go func(mychannel chan net.Conn) {
		for {
			select {
			case client := <-mychannel:
				//do nothing
				go deal(client)
			//default:
			//	fmt.Errorf("TASK_CHANNEL is full!")
			}
		}
	}(CONNECT_CHANNEL)

	for {
		client, err := listener.Accept()
		if client == nil {
			fmt.Printf("couldn't accept: ", err.Error()) // + err.Errors())
			continue
		}
		i++
		fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())

		select {
		case CONNECT_CHANNEL <- client:
			//do nothing
		default:
			//warnning!
			client.Close()
			fmt.Errorf("TASK_CHANNEL is full!")
		}
	}
}

/*
设计思路

https://studygolang.com/articles/2423
*/