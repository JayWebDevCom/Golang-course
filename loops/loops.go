package main

import "fmt"

func main() {
	// myLoop()
	// myLoopBreak()
	// myFor()
	myContinue()
}

func myLoop() {
	for i := 0; i <= 5; i++ {
		for j := 5; j >= 0; j-- {
			print(i, j)
		}
	}
}

func myLoopBreak() {
	i := 0
	for i <= 5 {
		if i == 3 {
			break
		}
		i++
	}
	print2(i)
	fmt.Println("loop over")
}

func myFor() {
	i := 0
	for i < 5 {
		fmt.Println("i is", i)
		i++
	}
}

func myContinue() {
	// will only print odds
	for i := 0; i <= 10; i++ {
		if i%2 != 1 {
			continue
		}
		fmt.Println("i is", i)
	}
}

func print2(int1 int) {
	fmt.Println("i is", int1)
}

func print(int1 int, int2 int) {
	fmt.Println("i", "is", int1, "j", "is", int2)
}
