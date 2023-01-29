package main

import (
	"fmt"
	"reflect"
)

// declaring a person struct
type Person struct {

	// setting the default value
	// of name to "geek"
	name string `default:"geek" default2:"12"`
}

func default_tag(p Person) string {

	// TypeOf returns type of
	// interface value passed to it
	typ := reflect.TypeOf(p)
	fmt.Printf("Type of p %v\n", typ)

	// checking if null string
	if p.name == "" {

		// returns the struct field
		// with the given parameter "name"

		f, _ := typ.FieldByName("name")

		// returns the value associated
		// with key in the tag string
		// and returns empty string if
		// no such key in tag
		p.name = f.Tag.Get("default2")
	}

	return fmt.Sprintf("%s", p.name)
}

// main function
func main() {

	// prints out the default name
	fmt.Println("Default name is:", default_tag(Person{}))
}
