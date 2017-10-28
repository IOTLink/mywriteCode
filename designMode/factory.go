package main


import (
	"fmt"
)

type Operater interface { //运算接口
	Operate(int, int) int
}

type AddOperate struct {  //加法运算类，实现了运算接口
}

func (this *AddOperate)Operate(rhs int, lhs int)  int {
	return rhs + lhs
}

type SubOperate struct { //减法运算类，实现了运算接口
}

func (this *SubOperate)Operate(rhs int, lhs int) int {
	return rhs - lhs
}

type OperateFactory struct {
}

func NewOperateFactory()  *OperateFactory {
	return &OperateFactory{}
}

func (this *OperateFactory)CreateOperate(operatename string) Operater {
	switch   operatename {
	case "+":
		return &AddOperate{}
	case "-":
		return &SubOperate{}
	default:
		panic("无效运算符号")
		return nil
	}
}


func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Operator := NewOperateFactory().CreateOperate("+")  //这时候Operator 不可能为nil
	fmt.Printf("add result is %d\n", Operator.Operate(1,2)) //add result is 3
	return
}

/*
简单工厂
*/