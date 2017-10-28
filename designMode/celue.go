package main
import (
	"fmt"
	"reflect"
)
type Compositer interface {
	SetCompositer(composition interface{})
	Compose()
}
type Composition struct{}
func (com *Composition) SetCompositer(composition interface{}) {
	fmt.Println("Here is seting ", reflect.TypeOf(composition), "before formattig")
}
func (com *Composition) Compose() {
	fmt.Println("Here is formatting")
}
func NewCompositor(strategy string) Compositer {
	composition := new(Composition)
	var g Compositer
	switch strategy {
	case "Array":
		g = &ArrayCompositor{composition}
	case "Tex":
		g = &TexCompositor{composition}
	case "Simple":
		g = &SimpleCompositor{composition}
	}
	return g
}
type ArrayCompositor struct {
	*Composition
}
type TexCompositor struct {
	*Composition
}
type SimpleCompositor struct {
	*Composition
}
func main() {
	c1 := NewCompositor("Array")
	c1.SetCompositer(c1)
	c1.Compose()
}

/*
https://www.golangtc.com/t/57744181b09ecc02f70001a3
策略模式
*/