package main

import "fmt"

// only one parameter type need be declared if the
// two are the same?
func greet(fname, lname string) string {
	// return fmt.Sprintf("Hello :) %s %s", fname, lname)

	// String print = Sprint
	return fmt.Sprint("Hello :) ", fname, lname)
}

func greetAndReturnVar(firstName string, lastName string) (s string) {
	s = fmt.Sprint("Hello again :) ", firstName, lastName)
	return
}

func greetAndReturns(firstName string, lastName string) (string, string) {
	return fmt.Sprint("Hello #3 ", firstName, lastName),
		fmt.Sprint("Hello #4 ", firstName, lastName)
}

func greetPeople(names ...string) (myString string) {
	for index, element := range names {
		myString += fmt.Sprintf("Hello %d %s ", index, element)
	}
	return
}

func average(numbers ...float64) (total float64) {
	for _, element := range numbers {
		total += element
	}
	total = total / float64(len(numbers))
	return
}
