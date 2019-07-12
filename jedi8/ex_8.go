package jedi8

import (
	"fmt"
)

// Ex8 fmt
func Ex8() {
	fmt.Println("Ex8")

	tyler := jack()
	fmt.Println(tyler())
}

func jack() func() string {
	return func() string {
		return fmt.Sprintln("Who am i?")
	}
}
