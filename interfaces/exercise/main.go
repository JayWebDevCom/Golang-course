package main

// https://godoc.org/sort
import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return 1
}

func (p people) Swap(a int, b int) {

}

func (p people) Less(a int, b int) bool {
	return a < b
}

type nums []int

// Len is exported
func (n nums) Len() int {
	return len(n)
}

// Less is exported
func (n nums) Less(a, b int) bool {
	fmt.Println("Less is called")
	return a < b
}

func (n nums) Swap(a, b int) {
	fmt.Println("Swap is called")
	var first int
	var second int

	for p, v := range n {
		fmt.Printf("p is %d, v is %d \n", p, v)
		if v == a {
			first = p
		}
		if v == b {
			second = p
		}
	}

	n[first] = b
	n[second] = a

}

func main() {
	// group := people{"a", "b", "c", "d", "e", "f", "g"}
	group := people{"h", "a", "b", "c", "d", "e", "f", "g"}
	sort.Slice(group, func(i, j int) bool {
		return group[i] < group[j]
	})

	fmt.Println(group)

	fmt.Println()

	numbers := nums{9, 8, 7, 6, 5, 4, 3}
	sort.Ints(numbers)
	fmt.Printf("first result is: %d\n", numbers[0])
	fmt.Println(numbers)

}
