package jedi3

import (
	"fmt"
)

// Ex3 fmt
func Ex3() {
	fmt.Println("Ex3")

	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("Who is Tyler Durder?")
	}()
	fmt.Println(42)
}

func bar() {
	fmt.Println(1984, "Big Brother is Wathincg ya")
}
