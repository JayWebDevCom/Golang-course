package main

import "fmt"

func getFunc() func(int) {
	return func(number int) {
		fmt.Printf("The number passed is %d\n", number)
	}
}

func visit(numbers []int, myCallback func(int)) {
	for _, number := range numbers {
		myCallback(number)
	}
}
