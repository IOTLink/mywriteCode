package main

import (
	"fmt"
)

type Pet struct {
	name string
}

type Dog struct {
	Pet
	Breed string
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed) //可以直接访问父类结构体的字段和方法。
}

func main() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
}

/*
http://zuozuohao.github.io/2016/06/16/Object-Oriented-Inheritance-in-Go/

liuhy@liuhy ~/work/src/mywriteCode/interface2 $ ./main
spot
my name is spot and I am a pointer

*/