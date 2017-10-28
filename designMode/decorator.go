package main

import (
	"fmt"
)

type Decoratorer interface {
	Draw()
}

type Composition struct {}

type Border struct {}

type Scroller struct {}

type BorderScrollerPage struct {
	c *Composition
	s *Scroller
	b *Border
}

func (c *Composition) Draw() {
	fmt.Println("Draw Composition")
}

func (b *Border) Draw() {
	fmt.Println("do something before drawwing border")
	b.DrawBorder()
}

func (b *Border) DrawBorder() {
	fmt.Println("Draw border")
}

func (s *Scroller) Draw() {
	fmt.Println("do something before drawwing scroller")
	s.DrawScroller()
}

func (s *Scroller) DrawScroller() {
	fmt.Println("Draw scroller")
}

func (bs *BorderScrollerPage) Draw() {
	bs.c.Draw()
	bs.s.Draw()
	bs.b.Draw()
	fmt.Println("Complete BorderScrollerPage")
}

func main() {
	bs := &BorderScrollerPage{&Composition{}, &Scroller{}, &Border{}}
	bs.Draw()
}

/*
修饰模式

http://zuozuohao.github.io/2016/07/06/Golang-Design-Patterns-Decorator/

Draw Composition
do something before drawwing scroller
Draw scroller
do something before drawwing border
Draw border
Complete BorderScrollerPage


 */