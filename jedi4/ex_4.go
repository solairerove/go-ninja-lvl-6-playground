package jedi4

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last)
}

// Ex4 fmt
func Ex4() {
	fmt.Println("Ex4")

	p := person{
		first: "Tyler",
		last:  "Durden",
		age:   23,
	}

	p.speak()
}
