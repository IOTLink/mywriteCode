package  main

import (
	"fmt"

)

type Node struct {
	pre  *Node
	data int
	next *Node

}
var head *Node
var tail *Node

func init() {
	head = &Node{nil,100,nil}
	tail = head
}

func InitItem(count int) {
	for i := 0; i< count; i++ {
		item := &Node{nil,i,nil}

		if i != 0 {
			hnext := head.next
			head.next = item
			item.pre = head
			item.next = hnext
			hnext.pre = item
		} else {
			hnext := head.next
			head.next = item
			item.pre = head
			item.next = hnext
			tail = item
			//fmt.Println("tial -> ",tail)
		}

		//fmt.Println(item)
	}
	//fmt.Println(tail)
}

func Query(begin int,end int) {
	p := head
	for i := begin; i< end-begin; i++ {
		//fmt.Println(p, p.data)
		p = p.next
	}
}

func Querytail() {
	p := tail
	 for p != nil {
		fmt.Println(p, p.data)
		p = p.pre
		//fmt.Println(p, p.data)
	}
}

func main() {
	InitItem(100)
	Query(0,101)
	Querytail()
}
