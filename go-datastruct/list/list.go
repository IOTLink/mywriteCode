package main

import (
	"fmt"
)

type Node struct {
	pre *Node
	next *Node
	value int64
}

type List struct {
	Head  *Node
	Tail  *Node
	Count int64
}

func CreateNode() (*List, error) {
	listHead := &List{nil,nil, 0}
	return listHead, nil
}

func IsEmpty(head *List) bool {
	isEmpty := false
	if head == nil {
		fmt.Println("head is nil")
	}
	if head.Count == 0 {
		isEmpty = true
	}
	return isEmpty
}

func AddHeadNode(head *List, item *Node) (*List, error) {
	var tempHead *List
	var err error
	if head == nil {
		tempHead, err = CreateNode()
		if err != nil {
			return nil,err
		}
	} else {
		tempHead = head
	}

	if tempHead.Count == 0 {
		tempHead.Head = item
		tempHead.Tail = item
		item.pre = nil
		item.next = nil
		tempHead.Count = 1
	} else {
		temp := tempHead.Head
		tempHead.Head = item
		item.pre = nil
		item.next = temp
		temp.pre = item
		tempHead.Count += 1
	}
	return tempHead, nil
}

func AddTailNode(head *List, item *Node) (*List, error) {
	if head == nil  {
		fmt.Println("head is nil")
	}

	if IsEmpty(head) {
		fmt.Println("head is empty")
		head.Head = item
		head.Tail = item
		head.Count += 1
		return head, nil
	}
	tail := head.Tail

	tail.next = item
	item.pre = tail
	item.next = nil

	head.Tail = item

	head.Count += 1
	return head,nil
}

func DelNode(head *List, n int64) bool {
	var count int64
	if head == nil {
		fmt.Println("head is nil")
	}

	if IsEmpty(head) {
		fmt.Println("head is empty")
		return false
	}

	if head.Count/2 - n >= 0 {
		for item := head.Head; item != nil; item = item.next {
			count += 1
			if count == n {
				tempPre  := item.pre
				tempNext := item.next
				tempPre.next = tempNext
				tempNext.pre = tempPre
				head.Count -= 1
				fmt.Println("del", item.value)
			}
		}
	} else {
		countTail := head.Count - n + 1
		for item := head.Tail; item != nil; item = item.pre {
			count += 1
			if count == countTail {
				tempPre  := item.pre
				tempNext := item.next
				tempPre.next = tempNext
				tempNext.pre = tempPre
				head.Count -= 1
				fmt.Println("del", item.value)
			}
		}
	}
	return true
}

func DisplayList(head *List) {
	if head == nil {
		fmt.Println("head is nil")
	}

	fmt.Println("List head ",head.Head.value)
	fmt.Println("List tail ",head.Tail.value)
	fmt.Println("List Count ",head.Count)

	for item := head.Head; item != nil; item = item.next {
		fmt.Println(item.value)
	}
}



func Find() {

}

func FindIndex() {

}

func main() {
	var err error
	head, _ := CreateNode()
	/*
	for i :=0; i<50; i++{
		item := &Node{nil,nil,int64(i)}
		head, err = AddHeadNode(head, item)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	*/
	for i := 0; i<100; i++{
		item := &Node{nil,nil,int64(i)}
		head, err = AddTailNode(head, item)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	DisplayList(head)
}