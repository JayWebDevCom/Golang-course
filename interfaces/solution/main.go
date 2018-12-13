package main

// https://godoc.org/sort
import (
	"fmt"
	"sort"
)

type people []string

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Len() int {
	return len(p)
}

type numbersz []int

func main() {
	group := people{"henrietta", "georgia", "fiona", "alexandra", "becky", "chanel", "deona", "erikah"}
	fmt.Println(group)
	fmt.Println("Now sorting...")
	sort.Sort(group)
	fmt.Println(sort.StringsAreSorted(group))
	fmt.Println(group)
	fmt.Printf("the type of group is: %T\n", group)

	fmt.Println()
	group = people{"henrietta", "georgia", "fiona", "alexandra", "becky", "chanel", "deona", "erikah"}
	fmt.Println("Now reversing...")
	sort.Sort(sort.Reverse(sort.StringSlice(group)))
	fmt.Println("reversed group: ")
	fmt.Println(group)

	fmt.Println()

	persons := []string{"Xa", "My", "Ta", "Co"}
	fmt.Println(persons)
	fmt.Println("Now sorting...")
	sort.Strings(persons)
	fmt.Println(sort.StringsAreSorted(persons))
	fmt.Println(persons)
	fmt.Println("Now reversing...")
	sort.Sort(sort.Reverse(sort.StringSlice(persons)))
	fmt.Println(persons)
	fmt.Printf("the type of persons is: %T\n", persons)

	fmt.Println()

	persons = []string{"Xa", "My", "Ta", "Co"}
	fmt.Println(persons)
	fmt.Println("Now sorting...")
	sort.StringSlice(persons).Sort()
	fmt.Println(persons)

	fmt.Println()

	persons = []string{"Xa", "My", "Ta", "Co"}
	fmt.Println(persons)
	fmt.Println("Now sorting...")
	sort.Sort(sort.StringSlice(persons))
	fmt.Println(persons)

	fmt.Println()

	numbers := []int{9, 8, 7, 6, 5, 4, 3}
	fmt.Println(numbers)
	fmt.Println("Now sorting...")
	sort.Ints(numbers)
	fmt.Printf("first result is: %d\n", numbers[0])
	fmt.Println(numbers)
	fmt.Printf("the type of numbers is: %T\n", numbers)

	fmt.Println()

	nums := numbersz{9, 8, 7, 6, 5, 4, 3}
	fmt.Println(nums)
	fmt.Println("Now sorting...")
	fmt.Println(sort.IsSorted(sort.IntSlice(nums)))
	sort.Ints(nums)
	fmt.Println(sort.IsSorted(sort.IntSlice(nums)))
	fmt.Printf("first result is: %d\n", nums[0])
	fmt.Println(nums)
	fmt.Printf("the type of numbers is: %T\n", nums)

	fmt.Println()

	nums = numbersz{9, 8, 7, 6, 5, 4, 3}
	fmt.Println(nums)
	x := sort.IntSlice(nums)
	fmt.Printf("the type of x is: %T\n", x)
	fmt.Println(sort.IsSorted(x))
	sort.Ints(x)
	fmt.Println(sort.IsSorted(sort.IntSlice(x)))
	fmt.Printf("first result is: %d\n", x[0])
	fmt.Println(x)
	fmt.Printf("the type of x is: %T\n", x)

}
