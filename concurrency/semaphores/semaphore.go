package semaphores

import "fmt"

/*
A semaphore is simply a variable.
This variable is used to solve critical section problems and to achieve process synchronization in the multi processing environment.
A trivial semaphore is a plain variable that is changed (for example, incremented or decremented, or toggled) depending on programmer-defined conditions
*/

// Semaphore is exported
func Semaphore() {

	/*
	many functions writing to one channel

	new channel takes a bool - put bool onto channel
	  */

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		// put true onto the channel
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		// put true onto the channel
		done <- true
	}()

	// waits until a value comes through
	go func() {
		// dispose of values - take off channel
		<-done
		<-done
		close(c)
	}()

	// blocks because it is an unbuffered channel
	for n := range c {
		fmt.Println(n)
	}
}

// SemaphoreLoop is exported
func SemaphoreLoop() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < n; i++ {
			go func() {
				for i := 0; i < 9; i++ {
					c <- i
				}
				done <- true
			}()
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
