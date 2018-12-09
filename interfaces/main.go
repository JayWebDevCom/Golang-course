package main

import (
	"fmt"
	"math"
)

/*
  in go interfaces are implemented
  no need to explicitly declare that 'implements the interface...'

	interfaces are types that just declare behaviour
*/

type shape interface {
	area() float64
	myInfo() string
}

func whatIsThis(i shape) {
	s := fmt.Sprintf("thae area of the interface implementation is %f", i.area())
	fmt.Println(s)
	fmt.Println(i.myInfo())
}

type square struct {
	side    float64
	details string
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z square) myInfo() string {
	return z.details
}

type circle struct {
	side    float64
	details string
}

func (z circle) area() float64 {
	return math.Pi * z.side * z.side
}

func (z circle) myInfo() string {
	return z.details
}

func info(z shape) {
	fmt.Println(z.myInfo())
	fmt.Println(z.area())
}

func main() {
	s := square{10, "I am a square"}
	c := circle{4, "I am a circle"}

	info(s)

	fmt.Println()

	info(c)
}
