package main

import "fmt"

// Assertion is used for interfaces
// syntax for assertion is .(type)

func main() {
	// name := "bella"
	// cant do this below because can't assert on non interface types
	// "bella" is a string
	// str, err := name.(string)
	// invalid type assertion: name.(string) (non-interface type string on left

	// if err {
	// 	fmt.Println("there was an error", err)
	// } else {
	// 	fmt.Println("%q\n", str)
	// }

	var name interface{} = "bella"
	str, ok := name.(string)

	if ok {
		fmt.Printf("%q\n", str)
		fmt.Printf("type is: %T", name)
	} else {
		fmt.Println("there was an error", ok)
		fmt.Println("value is not a string")
		fmt.Printf("type is: %T", name)
	}

	fmt.Println()

	var val interface{} = 5
	fmt.Printf("%T\n", val)

	// error mismatched types
	// nuances of the language
	// fmt.Printf(val + 4)
	fmt.Println(val.(int) + 4) // assertion not casting for interface types!

	fmt.Println()

}
