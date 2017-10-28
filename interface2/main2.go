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
	Token
	value uint64
}

func (m *Match) Type() TokenType {
	return m.toktype
}

func (m *Match) Lexeme() string {
	return m.lexeme
}
/*
func (i *IntegerConstant) Type() TokenType {
	return i.Token.Type()
}

func (i *IntegerConstant) Lexeme() string {
	return i.Token.Lexeme()
}

//./main2.go:58:12: cannot convert t (type IntegerConstant) to type Token:
//	IntegerConstant does not implement Token (Lexeme method has pointer receiver)


*/
func (i *IntegerConstant) Value() uint64 {
	return i.value
}

func main() {
	t := IntegerConstant{&Match{KEYWORD, "wizard"}, 2}
	fmt.Println(t.Type(), t.Lexeme(), t.Value())
	x := Token(t)
	fmt.Println(x.Type(), x.Lexeme())
}