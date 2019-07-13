package fanOutFanIn

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

/*
some channel ->

Fan out
 - multiple 'functions reading' from that channel until it is closed

Fan in
 - multiple 'channels writing' to that channel
 */

// FanOut is exported multiple 'functions reading' from that channel until it is closed
func FanOut() {
	c := fanIn(boring("Joe"), boring("Anne"))

	for i := 10; i > 0; i-- {
		// blocks because waiting to receive `(<-c)`
		// so program doesn't exit immediately
		fmt.Println(<-c)
	}

	fmt.Print("You are both boring I am leaving")
}

func fanIn(chan1, chan2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for { // infinite loop
			c <- <-chan1
		}
	}()

	go func() {
		for { // infinite loop
			c <- <-chan2
		}
	}()

	return c
}

func boring(message string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", message, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

// FanIn is exported
func FanIn() {
	channelOfInts := gen(2, 3, 4, 5)

	// Fan out
	// fan the work out
	// distributing the squaring work across two go
	// routine functions that both read from the same channel
	squares1 := square(channelOfInts)
	squares2 := square(channelOfInts)

	// Fan in
	// fan the results in
	// multiple 'channels writing' to that channel
	allSquares := merge(squares1, squares2)

	for e := range allSquares {
		fmt.Println(e)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for e := range in {
			out <- int(math.Pow(float64(e), float64(e)))
		}

		close(out)
	}()

	return out
}

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	// a channel to remove
	for _, channel := range channels {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(channel)
	}

	// a channel to wait and close that needs to wait for
	// the other go routine to finish
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
