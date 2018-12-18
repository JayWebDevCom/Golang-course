package main

import "fmt"
import "strconv"

// Conversion is used for converting one type to another e.g. int to float64
// Assertion is used for interfaces

func main() {
	// int to float is widening conversion
	var in int = 45
	var fl float64 = float64(in)
	fmt.Printf("converted to float64: %f\n", fl)

	// float to int is narrowing conversion
	var flo float64 = 23.99999
	var intt int32 = int32(flo)
	fmt.Printf("converted to int: %d\n", intt)

	// rune is an alias for int32 - unicode codepoint
	var x rune = 'a'
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)

	// convert rune to string
	fmt.Println("string of rune: ", string(x))
	fmt.Println("string of int32: ", string(y))

	// string to slice of bytes
	fmt.Println([]byte("hello world"))

	// https://golang.org/pkg/strconv/
	// i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	b, _ := strconv.ParseBool("true")

	//func ParseFloat(s string, bitSize int) (float64, error)
	// (multivalue return)
	f, _ := strconv.ParseFloat("3.1415", 64)

	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(s, b, f, i, u)

}
