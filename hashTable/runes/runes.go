package runes

import "fmt"

// RuneReminder is exported
func RuneReminder() {
	/*
			 'a' is an int32 because utf8 is a 1-4 byte coding scheme
			 i.e. seeing a string as a value
			 string is a collection of runes - runes are bytes
		   strings are slice of bytes
		   rune is alias for int32
	*/

	fmt.Printf("%v %T \n", 'a', 'a')
	fmt.Printf("%v %T \n", "a string", "a string")
	fmt.Printf("%v %T \n", "a string"[0], "a string"[0])
	fmt.Printf("%v %T \n", rune("a string"[0]), rune("a string"[0]))

	fmt.Println("Letter value at index 0:", rune("a string"[0]))
}
