package main

import (
	"fmt"
	"os"
)

const (
	indexFileNmae = "index.dat"
	fileName = "test.dat"
)

func InitFile() {
	initData := []byte("00000000")
	dataFile, err := os.Create(fileName)  //创建文件
	if err != nil {
		fmt.Println(err.Error())
	}
	var i uint64
	size := uint64(2<<20)
	for i=0; i<size; i++ {
		dataFile.Write(initData)
	}


	indexFile, err := os.Create(indexFileNmae)  //创建文件
	if err != nil {
		fmt.Println(err.Error())
	}

	return
	indexFile.WriteAt(initData,0)
	n, err := indexFile.WriteAt(initData,int64(size))
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(n)
		stat ,_ := indexFile.Stat()
		fmt.Println(stat.Size(),stat.Name())
	}
}

func writeFile() {
	file1, err := os.Create("test.txt")
	defer file1.Close()
	if err != nil {
	fmt.Println(file1, err)
	return
	}
	file1.WriteString("hello world")
}

func readFile() {
	file, err := os.Open("test.txt")
	defer file.Close()
	if err != nil {
	fmt.Println(file, err)
	}
	r_buf := make([]byte, 1024)

	for {
	n, _ := file.Read(r_buf)
	if n == 0 {
	break
	}
	fmt.Println(string(r_buf[0:n]))
	}
}

func main() {
	InitFile()
	//writeFile()
	//readFile()

}