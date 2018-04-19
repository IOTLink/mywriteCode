package main

import (
	"fmt"
)

type Singer interface {
	sing()
}

type Man struct {
	lyric string
}

type Bird struct {
	lyric string
}

func (m Man) sing() {
	fmt.Println(m.lyric)
}

func (b Bird) sing() {
	fmt.Println(b.lyric)
}

//go 实现的动态多态
func main() {
	var in Singer
	in = Man{"I'm a brave man"}
	in.sing()
	in = Bird{"I'm a small bird"}
	in.sing()
}

/*
interface　接口隐式继承

Ｃ＋＋　C++纯虚函数
https://blog.csdn.net/qq_36221862/article/details/61413619

Ｃ＋＋
静态多态
https://www.cnblogs.com/ivan0512/p/7701169.html
https://www.cnblogs.com/ivan0512/p/7701169.html


*/