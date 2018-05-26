package main

import "fmt"

func main() {
	fmt.Println(greet("Jon", "Doe"))
	fmt.Println(greetAndReturnVar("Jon", "Doe"))
	fmt.Println(greetAndReturns("Jon", "Doe"))
	fmt.Println(greetPeople("Louis", "Peter", "Paul"))

	fmt.Println(average(3, 33, 333, 3333))
}
