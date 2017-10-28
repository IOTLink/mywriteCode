package main

import (
	"os"
	//"io"
	"fmt"
	//"log"
)

type Datastore interface {
	AllBooks() ([]string, error)
}

type DB struct {
	fd  *os.File
}

func (db *DB)AllBooks() ([]string, error) {
	var i int
	arr := make([]string,10)
	for i= 0; i<10; i++{
		arr[i] = "string" + string(i)
	}
	return arr,nil
}


/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) (bool) {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func NewDB(filename string) (*DB, error) {
	var file *os.File
	var err error
	if checkFileIsExist(filename) {  //如果文件存在
		file, err = os.OpenFile(filename, os.O_APPEND, 0666)  //打开文件
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("文件存在");
	}else {
		file, err = os.Create(filename)  //创建文件
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("文件不存在")
	}
	return &DB{file},nil
}


type Env struct {
	db Datastore
}

func test(){
	var mydb DB
	mydb.fd = nil

	var db1 Datastore
	db1 = &mydb  //实例化的对象　赋值　　给抽象的虚类
	arr,_ :=db1.AllBooks()
	for i:=0; i<len(arr);i++{
		fmt.Println(arr[i])
	}

	switch v := db1.(type) {
	case Datastore:
		// v is an float64
		fmt.Printf("Datastore  %f\n",v)

	case *DB:
		// v is a string
		fmt.Printf("DB  %f\n",v)

	default:
		panic("I don't know how to handle this!")
	}

	db1 = db1.(*DB)
	db1.AllBooks()

}


func main() {

	/*
	db, err := NewDB("local.txt")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	arr,err := env.db.AllBooks()
	for i:=0; i<len(arr);i++{
		fmt.Println(arr[i])
	}
*/

	test()
}