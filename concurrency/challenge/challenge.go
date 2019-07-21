package challenge

import "fmt"

// DeadlockChallengeOne is exported
func DeadlockChallengeOne() {

	/*
		Deadlock
		c := make(chan int)
		c <- 1
		close(c)
		fmt.Println(<-c)
	*/

	// Fix
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()

	fmt.Println(<-c)
}

// DeadLockChallengeTwo is exported
func DeadLockChallengeTwo() {
	/*
		only prints zero
		c := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()

		fmt.Println(<- c)
	*/

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println("From Channel", n)
	}
}

// DeadLockChallengeThree is exported
func DeadLockChallengeThree() {
	/*
		all go routines are asleep
		c := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()

		for {
			fmt.Println(<-c)
		}
	*/

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for e := range c {
		fmt.Println(e)
	}
}

// DeadLockChallengeFactorial is exported
func DeadLockChallengeFactorial() {
	f := factorial(4)
	fmt.Println("Total factorial:", f)

	g := correctFactorialSolution(4)
	fmt.Println("Total factorial:", <-g)
}

func factorial(n int) int {

	/*
		factorial
		total := 1
		for i := n; i > 0; i-- {
				total += i
			}
		return total
	*/

	c := make(chan int)

	total := 1
	go func() {
		for i := n; i > 0; i-- {
			c <- i
		}
		close(c)
	}()

	// better to update total in the channel as #correctFactorialSolution shows
	for e := range c {
		total *= e
	}
	return total
}

func correctFactorialSolution(n int) chan int {
	c := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()

	return c
}
