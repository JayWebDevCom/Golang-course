package array

import . "fmt"

// Array1 is exported
func Array1() {

	var a [3]string = [3]string{"one", "two"}
	Println("-----")

	Printf("type: %T\n", a)
	Println("length: ", len(a))
	Println("first pos", a[1])

	// asci
	a[2] = string(45)

	for i := 0; i < len(a); i++ {
		Printf("position %d: %s\n", i, a[i])
	}

}

// Array2 is exported
func Array2() {
	Println("-----")
	var x [25]int
	Printf("length: %d\n", len(x))
	Printf("position 12 %d\n", x[12])

	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	for i, e := range x {
		Printf("%v - %T -  %b\n", e, e, e)
		if i == 10 {
			break
		}
	}

}
