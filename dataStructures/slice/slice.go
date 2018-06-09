package slice

import "fmt"

// Slice1 is exported
func Slice1() {
	fmt.Println("-----------")
	mySlice := []string{"a", "b", "c", "d"}
	fmt.Printf("Type: %T:\n", mySlice)
	fmt.Println("length:", len(mySlice))
	fmt.Println("slice is:", mySlice)
	fmt.Println("slicing:", mySlice[1:3])
}

// Slice2 is exported
func Slice2() {
	fmt.Println("-----------")
	mySlice := make([]int, 4, 4)
	fmt.Println("slice is", mySlice)

	mySlice2 := make([]int, 4, 4)
	// if only specify length - capacity will be same as length
	// append doubles length and capacity
	// appending to increace l or c can incur a performance cost
	fmt.Printf("Length before append: %v Capacity before append: %v\n", len(mySlice2), cap(mySlice2))
	mySlice2 = append(mySlice2, 4, 5, 6, 7)
	fmt.Printf("Length after append: %v Capacity after append: %v\n", len(mySlice2), cap(mySlice2))
	fmt.Println("index 6 in an array of previos length 4 but appended:", mySlice2[6])
}

// Slice3 is exported
func Slice3() {
	fmt.Println("-----------")

	// a slice of a slice of int
	// multidimentional slice
	sofs := [][]int{[]int{0, 1}, []int{0, 1}, []int{0, 1}}
	fmt.Println("slice of slice", sofs[2][1])

	sofs[2] = append(sofs[2], 2, 3, 4, 5, 6)
	fmt.Println("length of index 2:", len(sofs[2]))
	fmt.Println("slice of strings", sofs)

	// different types - intially of nil type
	var diff1 []interface{} = []interface{}{0, 1, true, 34.45, "hello world"}
	fmt.Println("different types", diff1)

	diff2 := make([]interface{}, 4)
	diff2[0] = "here"
	diff2[2] = 23456
	fmt.Println(diff2)
	fmt.Printf("Initial type: %v\n", diff2[1])
}

// Slice4 is exported
func Slice4() {

	slice := make([]string, 0)
	slice = append(slice, "one", "two", "three", "four", "five", "six", "seven")

	for i, e := range slice {
		fmt.Println("index", i, e)
	}

	fmt.Println("[0:3]")
	fmt.Println(slice[0:3])
	fmt.Println("[1:2]")
	fmt.Println(slice[1:2])
	fmt.Println("[1:3]")
	fmt.Println(slice[1:3])
	fmt.Println("[:2]")
	fmt.Println(slice[:2])
	fmt.Println("[5:]")
	fmt.Println(slice[5:])
	fmt.Println("[:]")
	fmt.Println(slice[:])
}

// Slice5 is exported - append a slice to a slice by unpacking args
func Slice5() {
	a := []float64{1.0, 1.1, 1.2}
	b := []float64{1.3, 1.4, 1.5}
	c := []float64{1.6, 1.7, 1.8}
	d := []float64{1.9, 2.0, 2.1}

	// cannot append more than one slice at a time
	e := append(a, append(b, append(c, d...)...)...)

	fmt.Println("e", e)
	fmt.Println("length of new slice", len(e))
}

// Slice6 is exported - deleting in a slice
func Slice6() {

	mySlice := []string{"one", "two", "three", "four", "five", "six", "seven"}
	fmt.Println("slice", mySlice)

	// variables can be used
	// to delete "three" (index 2) using a variable
	toDeleteIndex := 2
	mySlice = append(mySlice[:toDeleteIndex], mySlice[toDeleteIndex+1:]...)
	fmt.Println("slice", mySlice)
}

// Slice7 is exported - multidimentional slice
func Slice7() {
	records := make([][]string, 0)
	r1 := make([]string, 4)
	r1[0] = "name one"
	r1[1] = "a"
	r1[2] = "b"
	r1[3] = "c"

	r2 := make([]string, 4)
	r2[0] = "name two"
	r2[1] = "d"
	r2[2] = "e"
	r2[3] = "f"
	records = append(records, r1, r2)
	fmt.Println(records)

	greetings := make([]string, 3, 5)
	greetings = append(greetings, "hi")
	fmt.Println("greetings:", greetings)
	fmt.Println("greetings[0]:", greetings[0])
}

// Slice8 is exported - declaration
func Slice8() {

	fmt.Println("Declarations ")
	var x = []string{}
	fmt.Println("x:", x)
	//then must use append
	x = append(x, "first 1")
	fmt.Println("x:", x)

	var s []string
	fmt.Println("s:", s)
	//then must use append
	s = append(s, "first 2")
	fmt.Println("s:", s)

	p := make([]string, 5)
	// then can access indexes directly
	fmt.Println("length:", len(p), "capacity:", cap(p))
	fmt.Println("p:", p)
	p[3] = "three"
	fmt.Println("p:", p)
}
