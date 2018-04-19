package main

import (
	"fmt"
)

type Man struct {
	name string
	age  int
}


func (m Man) String() (result string) {
	result = fmt.Sprintf("I'm a man. My name is %s and I'm %d years old.\n", m.name, m.age)
	return
}



func main() {
	man := Man{"Bob", 18} //{Bob 18}

	fmt.Println(man)
}

/*
I'm a man. My name is Bob and I'm 18 years old.


*/