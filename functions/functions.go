package main

import "fmt"

// only one parameter type need be declared if the two are the same
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
	fmt.Println("Numbers is:", numbers)
	fmt.Printf("Type is %T \n", numbers)
	for _, element := range numbers {
		total += element
	}
	total = total / float64(len(numbers))
	return
}

func variadics() float64 {
	data := []float64{1, 2, 3, 57, 454, 8, 907, 3, 23, 342}
	// ... provides the slice as variadics i.e. pull each argument out one at a time
	ave := average(data...)
	return ave
}

func anotherAverage(numbers []float64) float64 {
	var total float64
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func variadicFunction() float64 {
	data := []float64{1, 2, 3, 57, 454, 8, 907, 3, 23, 342}
	return anotherAverage(data)
}
