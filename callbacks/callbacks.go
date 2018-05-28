package main

import "fmt"

// pass a function as a parameter to be called in the initial function
// calling back to the function passed in

func main() {
	data := []int{2, 3, 4}

	myFunc := getFunc()
	visit(data, myFunc)

	visit(data, func(number int) {
		fmt.Printf("The number passed in this anonymous function is %d\n", number)
	})

}
