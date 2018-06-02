package main

import "fmt"

// go has pass by value NOT pass by reference
// passing by value
// passing a memory address bypasses scope constraint
// passing the value of the memory address

func main() {
	age := 21
	changeMe(&age)
	fmt.Printf("age is now: %d \n", age)

	s := []string{"a", "b"}
	//s := make([]string, 1, 25)
	changeMe2(s)
	fmt.Printf("slice: %v \n", s)

	// reference types - ?
	myMap := make(map[string]int)
	changeMe3(myMap)
	fmt.Printf("map: %v \n", myMap)
}

func changeMe(x *int) {
	fmt.Printf("received: %d \n", x)
	*x = 12 // dereference
}

func changeMe2(z []string) {
	z[0] = "new"
}

func changeMe3(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
}
