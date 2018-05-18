package main

import "fmt"

func main() {

	a := 43

	// is assigned with the value of a's memory address
	// type is a pointer to an int
	var b *int = &a

	// * means de-reference a memory address to get the stored value
	fmt.Println(a, "memory address:", b, "dereferenced memory address:", *b)

	// the value at the address pointed to by b - change it to 14
	*b = 14

	// get the int at the memory address of the pointer b
	var c int = *b
	fmt.Println(a, c)

	// de-referencing a memory address of a variable and setting
	// it to something else can escape scope limitations
}
