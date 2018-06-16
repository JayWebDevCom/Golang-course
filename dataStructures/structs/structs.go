package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) fullName() {
	fmt.Println(p.firstName + " " + p.lastName)
}

func (p person) greet(food string) {
	fmt.Printf("Hello from %s I like to eat %s\n", p.firstName, food)
}

func (p person) sleep() {
	fmt.Println("I am now sleeping")
}

// reusability in inhericance
// boss has person fields and methods 'promoted' to it - Promotion.
// outer type field values override the inner type field values

// inner type promotion
// outter type overriding
type boss struct {
	person
	designation string
}

func (b boss) exclaim(name string, message string, p person) {
	fmt.Printf("Addressed to %s, I am %s, %s\n", name, b.firstName, message)
	p.greet("different things")
}

// method override by outer type
func (b boss) sleep() {
	fmt.Println("I don't like sleeping")
}
