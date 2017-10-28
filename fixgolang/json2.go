package main


import (
	"encoding/json"
	"fmt"
)

type PersonFlexible struct {
	Name interface{}
}

type Person struct {
	Name string
}

/*
func main() {
	theJson := `{"Name": 123}`

	var personFlexible PersonFlexible
	json.Unmarshal([]byte(theJson), &personFlexible)

	if _, ok := personFlexible.Name.(string); !ok {
		panic("Name must be a string.")
	}

	// When validation passes we can use the real object and types.
	// This code will never be reached because the above will panic()...
	// But if you make the Name above a string it will run the following:
	var person Person
	json.Unmarshal([]byte(theJson), &person)

	fmt.Printf("%+v\n", person)
}
*/

func main() {
	theJson := `123`

	var anything interface{}
	json.Unmarshal([]byte(theJson), &anything)

	switch v := anything.(type) {
		case float64:
		// v is an float64
		fmt.Printf("NUMBER: %f\n", v)

		case string:
		// v is a string
		fmt.Printf("STRING: %s\n", v)

		default:
		panic("I don't know how to handle this!")
	}
}