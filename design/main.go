package  main

import "fmt"

type Node struct {
	pre  *Node
	next *Node
	data int
}
var head *Node

func init() {
	head = &Node{nil,nil,100}
}

func InitItem(count int) {
	for i := 0; i< count; i++ {
		item := &Node{nil,nil,i}
		hnext := head.next
		head.next = item
		item.pre = head
		item.next = hnext
		//fmt.Println(item)
	}
}

func Query(begin int,end int) {
	p := head
	for i := begin; i< end-begin; i++ {
		fmt.Println(p.data)
		p = p.next
	}
}

func main() {
	InitItem(100)
	Query(0,100)

}
