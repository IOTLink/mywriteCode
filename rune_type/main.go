package main

import (
	"fmt"
)

func SubString(str string,begin,length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	fmt.Println("lth:",lth)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	fmt.Println("lth:",lth,rs[begin:end],string(rs[begin:end]))
	return string(rs[begin:end])
}


// rune:　Rune 是int32 的别名
//在UTF-8 世界的字符有时被称作runes。通常，当人们讨论字符时，多数是指8 位字符。UTF-8 字符可能会有32 位
func main() {
	s:="Go编程"
	fmt.Println(len(s))

	fmt.Println(len([]rune(s))) //个数

	s = "中华人民共和国"
	substr := SubString(s,2,2)
	fmt.Println(substr)
}
