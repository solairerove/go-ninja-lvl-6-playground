package jedi9

import (
	"fmt"
)

// Ex9 fmt
func Ex9() {
	fmt.Println("Ex9")

	foo(whoAmI, "Tyler", "Jack")
}

func whoAmI(name string) {
	switch name {
	case "Tyler":
		fmt.Println("I am Jack")
	case "Jack":
		fmt.Println("I am Tyler")
	default:
		fmt.Println("Poor Bob")
	}
}

func foo(f func(name string), names ...string) {
	for _, v := range names {
		f(v)
	}
}
