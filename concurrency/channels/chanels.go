package channels

import (
	"fmt"
	"time"
)

/*
https://golang.org/doc/effective_go.html#concurrency
dont use mutex, atomicity - share the variable
channels allow communication between goroutines

slices maps and channels use make()
data structures that must be initialize before used
 */

// UnbufferedChannel is exported
func UnbufferedChannel() {
	/*
	bufferred channel would be c := make(chan int, 50)
	unbufferred channel
	 */
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// put number onto channel
			// code stops until the value is taken from the channel
			// like a relay race
			c <- i
		}
	}() // self executing anonymous function

	go func() {
		for i := 0; i < 10; i++ {
			// take the number off the channel
			// receive the value from the channel and print it
			v := <-c
			fmt.Println(v)

		}
	}()

	time.Sleep(time.Second)
}

// ClosedChannel is exported
func ClosedChannel() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// put number onto channel
			// code stops until the value is taken from the channel
			// like a relay race
			c <- i
		}
		/*
		close - can no longer put values on, but can receive
		 */
		close(c)
	}()

	for n := range c {
		/*
		loop and get and print until channel closed and empty
		won't start until all values are put onto the channel
		 */
		fmt.Println(n)
	}
}
