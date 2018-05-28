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

func getFunc2() func(int) bool {
	return func(n int) bool {
		return n%3 == 0
	}
}

func filter(numbers []int, myCallback func(int) bool) []int {
	xs := []int{}
	for _, number := range numbers {
		if myCallback(number) {
			xs = append(xs, number)
		}
	}
	return xs
}
