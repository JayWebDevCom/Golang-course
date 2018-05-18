package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	fmt.Println("Enter meters to be converted\n>>")
	var meters float64

	// to a specific memory address
	fmt.Scan(&meters)
	yards := meters * metersToYards

	fmt.Println(meters, " meters is ", yards, " yards")
}
