package main

import (
	"fmt"
	"time"

	"runtime"
	"sync"

	"./atomicity"
	"./channels"
	"./channelsAsArguments"
	"./mutexConcurrency"
	"./racecondition"
	"./sameChannel"
	"./semaphores"
)

/*
 concurrency with go-routines
 without go, wg - will have 3 threads of go-routines - main function, foo function, bar function

 concurrency - doing many things at once - but only one at a time
 paralellism - doing many things at the same time

 go run main.go
 go run -race main.go
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

	fmt.Println("Now demonstrating the mutex")
	mutexConcurrency.Mutex()

	fmt.Println("Now demonstrating the atomicity")
	atomicity.Atomicity()

	fmt.Println("Now demonstrating the unbuffered channel")
	channels.UnbufferedChannel()

	fmt.Println("Now demonstrating the closed channel")
	channels.ClosedChannel()

	fmt.Println("Now demonstrating functions writing to the same channel")
	sameChannel.SameChannel()

	fmt.Println("Now demonstrating semaphores")
	semaphores.Semaphore()

	fmt.Println("Now demonstrating semaphores in a loop")
	semaphores.SemaphoreLoop()

	fmt.Println("Now demonstrating channels as arguments")
	channelsAsArguments.ChannelsAsArguments()
}
