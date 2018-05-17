package main

import "fmt"

var myName = "Jaiye"
var theName = "Johnny"

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
