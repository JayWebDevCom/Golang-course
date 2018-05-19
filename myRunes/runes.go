package main

import "fmt"

// rune is an alias for int32
// utf-8 is a 4 byte coding scheme
// with int32 (also 4 bytes ie 32 bits) we have numbers for each of the code points
// an integer identifying a unique code point

func main() {
	fmt.Println([]byte("byte list of this string"))
}
