package main

// golang.org, golang-book.com, blog.golang.org/index

import "fmt"

func ok() {
	fmt.Println("ok 2")
}

func main() {
	fmt.Println("ok")
	ok()
	fmt.Println(myName)
	printString("Riddik")

	// assigning a function to a variable (from vars file)
	changeName := func() string {
		theName = "Papi is the new name "

		return theName
	}

	changeName()

	fmt.Println(theName)

	fmt.Println(wrapper()())

	// httpStuff()

	fmt.Println(myConstant)

	fmt.Println(aio, bio, cio, KB, MB)

	theIotas()

	memory()
}
