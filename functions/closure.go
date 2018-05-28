package main

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
