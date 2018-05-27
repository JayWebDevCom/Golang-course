package main

import "fmt"

// returns a func which returns a string
func makeGreeter() func() string {
	return func() string {
		return "Hello make greeter"
	}
}

func main() {

	// assigning a function to a variable is a func expression
	nowGreet := func() {
		fmt.Println("Hello world")
	}

	nowGreet()
	fmt.Printf("nowGreet is a %T \n", nowGreet)

	mGreeter := makeGreeter()
	fmt.Printf("mGreeter is a %T \n", mGreeter)
	fmt.Println(mGreeter())

	fmt.Println(greet("Jon", "Doe"))
	fmt.Println(greetAndReturnVar("Jon", "Doe"))

	fmt.Println()
	fmt.Println(greetAndReturns("Jon", "Doe"))
	fmt.Println(greetPeople("Louis", "Peter", "Paul"))

	fmt.Println()
	fmt.Println(average(3, 33, 333, 3333))
	fmt.Println(variadics())
	fmt.Println(variadicFunction())
}
