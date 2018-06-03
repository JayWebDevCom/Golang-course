package main

import (
	"fmt"

	"./exercises"
)

// expressions that evaluate to a value of type bool!
// expressions happen horizontally
// statements happen vertically

// an expression produces a value and can be written wherever a value is expected
// a statement performs an action

func main() {
	if true || false {
		fmt.Println("This is true or false")
	}

	if true && false {
		fmt.Println("This will not run")
	}

	fmt.Println(exercises.Exercise1(15))

}
