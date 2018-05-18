package main

import "fmt"

var myName = "Jaiye"
var theName = "Johnny"
const myConstant string = "This does not change"

func printString(name string) {
	fmt.Println(name + " from a func defined in vars...")
}

// returning a function
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func memory() {
	a := "some memory info"
	output := fmt.Sprint(a, " memory address is: ", &a)
	fmt.Println(output)
}
