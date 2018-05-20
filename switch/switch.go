package main

import (
	"fmt"
	"reflect"
)

func main() {
	mySwitch("twoo")
	mySwitch2("three")

	fmt.Println(len("a very long string"))
	fmt.Println(reflect.TypeOf("a very long string"))
	fmt.Println(fmt.Sprintf("%T", "a very long string"))

	if myVar := "anything"; 3 < 5 {
		// myVar is initialized into a very small scope
		// and the subsequent condition is evaluated
		fmt.Println("ok:", myVar)
	}

}

func mySwitch(num string) {
	switch num {
	case "one":
		fmt.Println(1)
	case "two", "twoo":
		fmt.Println(2)
	default: // default is not needed
		fmt.Println("default")
	}
}

func mySwitch2(num string) {
	switch {
	case num == "one":
		fmt.Println(1)
	case num == "two", num == "twoo":
		fmt.Println(2)
	case num == "three":
		fmt.Println(3)
	default: // default is not needed
		fmt.Println("default")
	}
}
