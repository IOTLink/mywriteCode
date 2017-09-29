package main

import (
	"crypto/subtle"
	"fmt"
)

var lessOrEqTests = []struct {
	x, y, result int
}{
	{0, 0, 1},
	{1, 0, 0},
	{0, 1, 1},
	{10, 20, 1},
	{20, 10, 0},
	{10, 10, 1},
}

func TestConstantTimeLessOrEq() {
	for i, test := range lessOrEqTests {
		result := subtle.ConstantTimeLessOrEq(test.x, test.y)
		if result != test.result {
			fmt.Println("#%d: %d <= %d gave %d, expected %d", i, test.x, test.y, result, test.result)
		}
	}
}
func main() {
	ret := subtle.ConstantTimeByteEq(1,1)
	fmt.Println(ret)
	ret = subtle.ConstantTimeByteEq(19,1)
	fmt.Println(ret)

	ret = subtle.ConstantTimeEq(1,1)
	fmt.Println(ret)
	ret = subtle.ConstantTimeEq(11,1)
	fmt.Println(ret)

	ret = subtle.ConstantTimeLessOrEq(1,2)
	fmt.Println(ret)
	ret = subtle.ConstantTimeLessOrEq(11,2)
	fmt.Println(ret)

	x := []byte("hello world")
	y := []byte("ooooooooooo")
	subtle.ConstantTimeCopy(1,y,x)
	fmt.Println(string(y))
	TestConstantTimeLessOrEq()

	ss := subtle.ConstantTimeSelect(1,111,222)
	fmt.Println(ss)
}
//https://www.kancloud.cn/wizardforcel/golang-stdlib-ref/121507
//http://blog.codeg.cn/2015/01/12/go-source-code-reading-ConstantTimeByteEq/
//