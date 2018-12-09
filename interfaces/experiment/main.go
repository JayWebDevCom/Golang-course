package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	myInfo() string
}

type circle struct {
	side    float64
	details string
}

func (c circle) area() float64 {
	return math.Pi * c.side * c.side
}

func (c circle) myInfo() string {
	return c.details
}

func (c circle) someFunc() string {
	return "circle some func"
}

type square struct {
	name   string
	length int
}

func (s square) area() float64 {
	return float64(s.length)
}

func (s square) myInfo() string {
	return s.name
}

func (s square) someFunc() string {
	return "square some func"
}

func takesShape(s shape) float64 {
	return s.area()
}

func main() {
	c := circle{side: float64(3.2), details: "my circle"}
	s := square{name: "my square", length: 5}

	var ret1 = takesShape(s)
	var ret2 = takesShape(c)

	fmt.Println(ret1)
	fmt.Println(ret2)

	fmt.Println(c.someFunc())
	fmt.Println(s.someFunc())
}
