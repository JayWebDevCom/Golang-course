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

	fmt.Println(exercises.Exercise2(15, func(num int) bool {
		return float64(num)/2 == 0
	}))

	myFunc := func(x int) bool {
		return float64(x)/2 == 0
	}

	fmt.Println(exercises.Exercise2(15, myFunc))

	data := []float64{7, 3, 6, 3, 8, 2, 4, 1, 2, 6, 8, 3, 7, 8, 4}

	fmt.Println(exercises.Exercise3(data...))

	fmt.Println(exercises.Exercise4(1))
	fmt.Println(exercises.Exercise4(1, 2))
	fmt.Println(exercises.Exercise4(data...))

}
