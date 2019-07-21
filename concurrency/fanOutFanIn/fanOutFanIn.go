package fanoutfanin

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

/*
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
	// multiple 'functions reading' from that channel until it is closed
	squares1 := square(channelOfInts)
	squares2 := square(channelOfInts)

	// Fan in
	// fan the results in
	// multiple 'channels writing' to that channel
	// - multiple 'channels writing' to that channel
	allSquares := Merge(squares1, squares2)

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

// Merge is exported
func Merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	// avoids deadlocks
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
	// starts after w.Add()
	// calls wg.Wait() ie begin go routines
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

var workerID int
var publisherID int

// FOutFInChallenge is exported - it does fan out - but doesn't fan in
func FOutFInChallenge() {
	// fan out - yes many diff functions read from same channel
	// fan in - converging many channels onto one channel - multiplexing - not here
	// no fan in here
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)

	time.Sleep(1 * time.Millisecond)
}

func workerProcess(in chan string) {
	workerID++
	thisID := workerID

	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is : %s\n", thisID, input)
	}
}

func publisher(in chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0

	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", dataID)
		data := fmt.Sprintf("data from publisher %d: data: %d", thisID, dataID)
		in <- data
	}
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

// FanFactorial is exported
func FanFactorial() {

	// prepare numbers for factorial
	inNumbers := generateNumbers()

	// Fan out - multiple 'functions reading' from that channel until it is closed
	factorialResult1 := toFactorial(inNumbers)
	factorialResult2 := toFactorial(inNumbers)
	factorialResult3 := toFactorial(inNumbers)

	allResults := mergeChallenge(factorialResult1, factorialResult2, factorialResult3)

	for i := range allResults {
		fmt.Printf("Number is %d\n", i)
	}
}

// FanFactorialTwo is exported
func FanFactorialTwo() {

	// prepare numbers for factorial
	inNumbers := generateNumbers()

	// Fan out - multiple 'functions reading' from that channel until it is closed
	var chanSlice [] <-chan int

	for i := 0; i < 10; i++ {
		chanSlice = append(chanSlice, toFactorial(inNumbers))
	}

	fmt.Printf("There have been %d channels fanned out to\n", len(chanSlice))

	allResults := mergeChallenge(chanSlice...)

	counter := 0
	for result := range allResults {
		counter++
		fmt.Printf("This is iteration %d \t Number is %d\n", counter, result)
	}
}



func generateNumbers() <-chan int {
	intSlice := []int{10, 11, 12, 13, 14, 15, 16, 19, 18, 19}
	out := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			for _, j := range intSlice {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func toFactorial(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()

	return out
}

func mergeChallenge(inChannels ...<-chan int) <-chan int {

	// avoids deadlocks
	var wg sync.WaitGroup
	wg.Add(len(inChannels))

	out := make(chan int)

	for _, channel := range inChannels {
		go func(ch <-chan int) {
			for e := range ch {
				out <- e
			}
			wg.Done()
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
