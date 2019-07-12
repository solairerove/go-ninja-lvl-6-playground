package jedi

import (
	"fmt"
)

// Ex1 fmt
func Ex1() {
	fmt.Println("Ex1")

	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}
