package main

import "fmt"

// struct is a composed / aggregate type
// in go you create a type not a class and then
// create variables of that type
// struct is an aggregate type - a collection of
// other types encapsulated into one type

func main() {

	fmt.Println()
	jen := person{"Jennifer", "Bond", 34}
	fmt.Printf("Jennifer's age is %d\n", jen.age)
	jen.fullName()
	jen.greet("cake")
	jen.sleep()

	fmt.Println()
	marco := boss{
		// person struct literal
		person{
			"Marco",
			"Polo",
			40,
		},
		"programmer",
	}

	marco.greet("beans")

	// method overridden by outer type
	marco.sleep()
	// inner method overridden but still available
	marco.person.sleep()
	marco.exclaim("Johnny", "where are you?", person{"Mikey", "Jones", 20})

	fmt.Println()
	// assigning the address
	p1 := &person{"Oberyn", "Martell", 40}
	// printing shows that type is a pointer in the console
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	// pointers still make values and methods available
	p1.fullName()
	fmt.Println(p1.age)
}
