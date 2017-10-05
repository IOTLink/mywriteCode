package main


import (
	"bufio"
	"fmt"
	"net"
	"time"
	"strconv"
)

//server

const PORT = 7051

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn,10000)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: ",err.Error())// + err.Errors())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}


var quitSemaphore chan bool

func main() {
	//server
	go func() {
		// server, err := net.Listen("tcp", "192.168.1.106:" + strconv.Itoa(PORT))
		server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
		if server == nil {
			fmt.Printf("couldn't start listening: %s",err.Error())// + err.String())
			return
		}
		conns := clientConns(server)
		for {
			go handleConn(<-conns)
		}
	}()

	var tcpAddr *net.TCPAddr
	//localaddr := &net.TCPAddr{IP:net.ParseIP("192.168.1.107")}
	//localaddr := &net.TCPAddr{IP:net.ParseIP("192.168.1.107")}
	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.1.106:3333")
	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:3333")

	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "47.92.106.133:7051")
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "123.207.15.221:7051")

	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "58.31.231.231:7051")

	//conn, _ := net.DialTCP("tcp", localaddr, tcpAddr)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	//b := []byte("time\n")
	ts := time.Now().String()
	b := []byte("frist: "+ts + "\n" )
	conn.Write(b)

	<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		//ts := time.Now().String()
		//b := []byte("second: "+ts + "\n" )
		//conn.Write(b)
	}
}

