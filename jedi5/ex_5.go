package jedi5

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2.0)
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

// Ex5 fmt
func Ex5() {
	fmt.Println("Ex5")

	c := circle{
		radius: 20.20,
	}

	s := square{
		length: 10.00,
	}

	info(c)
	info(s)
}
