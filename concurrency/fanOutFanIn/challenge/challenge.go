package challenge

import (
	"fmt"
	"strconv"
)
import "../../fanoutfanin"

// DeadlockChallenge is exported
func DeadlockChallenge() {

	in := genDeadlockChallenge()

	// don't fanout in a go routine!
	chanSlice := fanOutDeadlockChallenge(in, 10)
	fmt.Printf("chanSlice arriving has %d entries\n", len(chanSlice))

	for n := range fanoutfanin.Merge(chanSlice...) {
		fmt.Println(n)
	}
}

func genDeadlockChallenge() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				out <- j
			}
		}
		close(out)
	}()

	fmt.Printf("out has %d entries\n", len(out))

	return out
}

// don't fanout in a go routine!
func fanOutDeadlockChallenge(in <-chan int, num int) [] <-chan int {
	var chanSlice []<-chan int

	for i := 0; i < num; i++ {
		chanSlice = append(chanSlice, factorialDeadlockChallenge(in))
	}

	fmt.Printf("chanSlice leaving has %d entries\n", len(chanSlice))

	return chanSlice
}

func factorialDeadlockChallenge(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for j := range in {
			out <- factDeadlockChallenge(j)

		}
		close(out)
	}()

	return out
}

func factDeadlockChallenge(num int) int {
	total := 1
	for i := num; i > 0; i-- {
		total *= i
	}
	return total
}

func IncrementorSolution() {
	in := incrementor(2)

	var count int
	for val := range in { // blocks because pulling from the channel
		count++
		fmt.Println(val)
	}

	fmt.Printf("Final count: %d\n", count)
}

func incrementor(num int) chan string {
	out := make(chan string)
	done := make(chan bool) // semaphore letting us know when finished and can close out channel

	for i := 0; i < num; i++ {
		go func(val int) {
			for j := 0; j < 20; j++ {
				// Itoa = Integer to asci
				out <- fmt.Sprint("Process: "+strconv.Itoa(val)+" printing:", j)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < num; i++ {
			<-done
		}
		close(out)
	}()

	return out
}
