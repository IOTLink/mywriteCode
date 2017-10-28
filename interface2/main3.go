package main

import (
	"fmt"
)


type TokenType uint16

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LBRACKET
	RBRACKET
	INT
)



type Token interface {
	Type()   TokenType
	Lexeme() string
}

type Match struct {
	toktype TokenType
	lexeme  string
}

type IntegerConstant struct {
	*Match
	value uint64
}

func (m *Match) Type() TokenType {
	return m.toktype
}

func (m *Match) Lexeme() string {
	return m.lexeme
}

func (i *IntegerConstant) Type() TokenType {
	return i.Match.Type()  //i.Type() 递归死掉
}

func (i *IntegerConstant) Lexeme() string {
	return i.Match.Lexeme()
}

func (i *IntegerConstant) Value() uint64 {
	return i.value
}

func main() {
	t := IntegerConstant{&Match{KEYWORD, "wizard"}, 2}
	fmt.Println(t.Type(), t.Lexeme(), t.Value())
	x := t.Match
	fmt.Println(x.Type(), x.Lexeme())
}
/*
崩溃

func (i *IntegerConstant) Type() TokenType {
	return i.Type()
}

func (i *IntegerConstant) Lexeme() string {
	return i.Lexeme()
}

liuhy@liuhy ~/work/src/mywriteCode/interface2 $ ./main
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow

runtime stack:
runtime.throw(0x4b9366, 0xe)
	/usr/local/go/src/runtime/panic.go:605 +0x95
runtime.newstack(0x0)
	/usr/local/go/src/runtime/stack.go:1050 +0x6e1
runtime.morestack()
	/usr/local/go/src/runtime/asm_amd64.s:415 +0x86

goroutine 1 [running]:
main.(*IntegerConstant).Type(0xc460083ed8, 0x0)
	/home/liuhy/work/src/mywriteCode/interface2/main3.go:43 +0x44 fp=0xc440084378 sp=0xc440084370 pc=0x487c04
main.(*IntegerConstant).Type(0xc460083ed8, 0x0)
	/home/liuhy/work/src/mywriteCode/interface2/main3.go:44 +0x2b fp=0xc440084398 sp=0xc440084378 pc=0x487beb
main.(*IntegerConstant).Type(0xc460083ed8, 0x0)
	/home/liuhy/work/src/mywriteCode/interface2/main3.go:44 +0x2b fp=0xc4400843b8 sp=0xc440084398 pc=0x487beb

 */