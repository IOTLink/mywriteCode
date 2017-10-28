package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	theJson := `{"FirstName": "Bob", "LastName": "Smith"}`

	var person Person
	json.Unmarshal([]byte(theJson), &person)

	fmt.Printf("%+v\n", person)



	person1 := Person{"James", "Bond"}
	theJson1, _ := json.Marshal(person1)

	fmt.Printf("%+v\n", string(theJson1))
}
