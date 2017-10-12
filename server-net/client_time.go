package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"encoding/json"
	"time"
	"strings"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "9999", "port")

type Msg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type Resp struct {
	Data string `json:"data"`
	Status int  `json:"status"`
}

type ServerCfg struct {
	IPv4 string
	Port string
}

func main() {
	/*
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	*/
	cfg := make([]ServerCfg, 4)
	if cfg == nil {
		fmt.Println("malloc abort!")
	}
	for i := 0; i < 3; i++{
		cfg[i].IPv4 = "12.0.0.1"
		cfg[i].Port = strconv.Itoa( 1025 + i)
	}
	cfg[3].IPv4 = "127.0.0.1"
	cfg[3].Port = strconv.Itoa( 9999)

	var conn net.Conn
	var err error
	var waitime time.Duration
	connecd := false

	for i := 0; i < 4; i++ {
		fmt.Println("Connecting to " + cfg[i].IPv4+":"+cfg[i].Port)
		for n := 0; n < 3; n++ {
			waitime = time.Second * (time.Duration(n) + 1)
			conn, err = net.DialTimeout("tcp", cfg[i].IPv4+":"+cfg[i].Port, waitime)
			if err != nil {
				if strings.Contains( err.Error(), "timeout") {
					fmt.Println("time out ", n+1, "second")
					continue
				} else {
					fmt.Println("Error connecting:",err.Error())
					os.Exit(1)
				}
			}
			connecd = true
			break
		}
	}

	if !connecd {
		fmt.Println("Not connect server")
		os.Exit(1)
	}
	// 下面进行读写
	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	// write 10 条数据
	for i := 10; i > 0; i-- {
		d := "hello " + strconv.Itoa(i)
		msg := Msg{
			Data: d,
			Type: 1,
		}
		// 序列化数据
		b, _ := json.Marshal(msg)
		writer := bufio.NewWriter(conn)
		_, e := writer.Write(b)
		//_, e := conn.Write(b)
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
		// 增加换行符导致server端可以readline
		//conn.Write([]byte("\n"))
		writer.Write([]byte("\n"))
		writer.Flush()
	}
	fmt.Println("Write Done!")
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	// 读取数据
	for i := 1; i <= 10; i++ {
		//line, err := reader.ReadString(byte('\n'))
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		// 反序列化数据
		var resp Resp
		json.Unmarshal(line, &resp)
		fmt.Println("Status: ", resp.Status, " Content: ", resp.Data)
	}
	fmt.Println("Read Done!")
}

