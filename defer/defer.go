package main

func main() {
	defer printThis("Hello")
	printThis(" there")
	printThis(" world!")

	// useful for file close at the end of an operation
}
