package pipelines

import (
	"fmt"
	"math"
	"sync"
)

/*
https://blog.golang.org/pipelines
 */

// Pipelines is exported
func Pipelines() {
	var generatedChannel <-chan int = generateChannel(2, 3, 4, 5, 6)

	// cube the cubes
	var actedOnValuesChannel <-chan int = powerOf(generatedChannel)
	var actedOnValuesChannel2 <-chan int = powerOf(actedOnValuesChannel)

	// range removes the channel
	for n := range actedOnValuesChannel2 {
		fmt.Println(n)
	}
}

// FactorialChallenge is exported is exported
func FactorialChallengeSolution() {

	// make channel of subject numbers
	channelOfNumbersToOperateOn := generateSubjectNumbers()

	// make channel of answer (factorial) numbers
	channelOfFactorials := generateFactorialNumbers(channelOfNumbersToOperateOn)

	for n := range channelOfFactorials {
		fmt.Println(n)
	}
}

func generateSubjectNumbers() <-chan int {
	out := make(chan int)

	// 2 loops to keep factorial number small
	// out has 100 nums to perform factorial operation on
	go func() {
		for i := 0; i < 10; i++ {
			// 10 times - put 1 to 10
			for j := 1; j < 11; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func generateFactorialNumbers(in <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for n := range in {
			c <- factorial(n)
		}
		close(c)
	}()

	return c
}

func factorial(num int) int {
	var total int = 1

	for i := num; i > 0; i-- {
		total *= i
	}

	return total
}

// FactorialChallenge is exported
func FactorialChallenge() {

	preparedChannel := prepareChannel()

	factorialChannel := factTheChannel(preparedChannel)

	for n := range factorialChannel {
		fmt.Println(n)
	}
}

func prepareChannel() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 1; j < 11; j++ {
				c <- j
			}
		}
		close(c)
	}()

	return c
}

func factTheChannel(in chan int) chan int  {
	channel := make(chan int)

	go func() {
		for n := range in {
			channel <- factorial(n)
		}
		close(channel)
	}()

	return channel
}


var wg sync.WaitGroup
var mutex sync.Mutex

// FactorialChallengeWithWGMutex is exported
func FactorialChallengeWithWGMutex() {
	var numIterations int = 2

	channelSlice := make([]chan int, 100)

	for i := 0; i < numIterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go func() {
				for j := 0; j < 10; j++ {

					factorialChannel := make(chan int)
					num := 1
					nextNum := num + j

					factorialChannel <- nextNum
					num = nextNum

					mutex.Lock()
					channelSlice = append(channelSlice, factorialChannel)
					mutex.Unlock()
					close(factorialChannel)
				}
			}()
		}()
	}

	wg.Wait()

	go func() {
		out := make(chan int)

		for e := range channelSlice {
			out <- e
		}

		close(out)
	}()

	for n := range channelSlice {
		fmt.Println(n)
	}
}

func generateChannel(nums ...int) <-chan int {
	var out chan int = make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func powerOf(inChannel <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range inChannel {
			// to power of 3
			pow := math.Pow(float64(n), float64(3))
			out <- int(pow)
		}
		close(out)
	}()

	return out
}
