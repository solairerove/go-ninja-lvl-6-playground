package jedi2

import (
	"fmt"
)

// Ex2 fmt
func Ex2() {
	fmt.Println("Ex2")

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func bar(xi []int) int {
	return foo(xi...)
}
