package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"fn"`
	LastName  string `json:"ln"`
}

func main() {
	theJson := `{"fn": "Bob", "ln": "Smith"}`

	var person Person
	json.Unmarshal([]byte(theJson), &person)

	fmt.Printf("%+v\n", person)
}
