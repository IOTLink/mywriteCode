package main

/*
组合模式

http://zuozuohao.github.io/2016/06/25/Golang-Design-Patterns-Composite/

*/



import (
	"fmt"
	"reflect"
)

type Appearancer interface {
	Draw(elemet Appearancer)
	Intersect(point int)
	SetParent(parentID int)
}

type Glyph struct {
	Name     string
	Position int
	ID       int //ID must > 0
	ParentID int //if ParentID equal 0, the Glyph has no parents
}

type Character struct {
	Glyph
}

type Rectangle struct {
	Glyph
}

type Row struct {
	Glyph
	Childs []Appearancer
}

func (g *Glyph) Draw(elemet Appearancer) {
	fmt.Println("I am a ", reflect.TypeOf(elemet), ":", g.Name)
}

func (g *Glyph) Intersect(point int) {
	if g.Position == point {
		fmt.Println(g.Name, " is far away from ", point)
	} else {
		fmt.Println(g.Name, " intersect with ", point)
	}
}

func (g *Glyph) SetParent(parentID int) {
	g.ParentID = parentID
}

func (r *Row) Insert(child Appearancer, position int) {
	index := r.insertInRightPlace(child, position)
	child.SetParent(r.ID)
	fmt.Println("Add ", child, "to Childs at position ", index)
	fmt.Println(r.Name, "'s length is ", len(r.Childs))
}

func (parent *Row) insertInRightPlace(child Appearancer, position int) int {
	insertedPosition := 0
	childsLength := len(parent.Childs)
	if position > (childsLength - 1) {
		parent.Childs = append(parent.Childs, child)
		insertedPosition = childsLength
	} else {
		parent.Childs = append(parent.Childs[position:position], child)
		insertedPosition = position
	}
	return insertedPosition
}

func main() {
	c1 := &Row{Glyph{"c1", 12, 1, 0}, []Appearancer{}}
	c1.Draw(c1)
	c1.Intersect(2)
	c1.Insert(&Character{Glyph{"c1", 12, 2, 0}}, 3)
	fmt.Println("hello Composite")
}
/*
I am a  *main.Row : c1
c1  intersect with  2
Add  &{{c1 12 2 1}} to Childs at position  0
c1 's length is  1
hello Composite

*/