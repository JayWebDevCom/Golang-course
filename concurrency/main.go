package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"./racecondition"
)

/*
 concurrency with go-routines
 without go, wg - will have 3 threads of go-routines - main function, foo function, bar function

 concurrency - doing many things at once - but only one at a time
 paralellism - doing many things at the same time
*/

func init() {
	// init is first code to execute
	// not necessary but this adds all cpu cores to the program
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}

func fooWait() {
	for i := 10; i < 15; i++ {
		fmt.Println("foo: ", i)
		// should introduce interleaving
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func barWait() {
	for i := 10; i < 15; i++ {
		fmt.Println("bar: ", i)
		// should introduce interleaving
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	wg.Add(2)
	go fooWait()
	go barWait()
	wg.Wait()

	fmt.Println("Now demonstrating the race condition")
	racecondition.Race()
}