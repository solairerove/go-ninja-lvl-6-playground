package jedi10

import (
	"fmt"
)

// Ex10 fmt
func Ex10() {
	fmt.Println("Ex10")

	tyler := foo()
	jack := foo()

	fmt.Println(tyler())
	fmt.Println(tyler())
	fmt.Println(jack())
}

func foo() func() string {
	s := "Jack"

	return func() string {
		if s == "Jack" {
			s = "Tyler"
		} else {
			s = "Jack"
		}
		return fmt.Sprintln(s)
	}
}
