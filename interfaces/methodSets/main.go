package main

import (
	"fmt"
	"math"
)

/*
  The method set of a type determines the interfaces that the type implements
  and the methods that can be called using a receiver of that type
*/

type shape interface {
	area() float64
	myInfo() string
}

type circle struct {
	side    float64
	details string
}

// here circle is a value receiver
func (z circle) area() float64 {
	return math.Pi * z.side * z.side
}

// here circle is a pointer receiver - will need to pass in an address when this method is called
func (z *circle) myInfo() string {
	return z.details
}

// my info is called which needs a pointer receiver
func info(z shape) {
	fmt.Println(z.myInfo())
}

func main() {
	c := circle{4, "I am a circle"}

	// info must be called with a pointer
	// value receiver can accept a value or an address
	// pointer receiver must take an address
	info(&c)
	fmt.Println()

	// Receivers -- values
	// (t T) can take T and *T
	// (t *T) can take only *T

	// stack is a downward growing history of method calls, eliminated as they return
	// heap is a store of variable escaping the stack
	// not all values have an address
	// no pointer receivers cant always be given a value because address may ont exist

}
