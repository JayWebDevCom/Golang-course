package channelsAsArguments

import (
	"fmt"
)

// ChannelsAsArguments is exported
func ChannelsAsArguments() {
	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		fmt.Println(n)
	}

}

// ChannelsDirection is exported
func ChannelsDirection() {
	cDir := incrementorDirection()
	cSumDir := pullerDirection(cDir)

	for n := range cSumDir {
		fmt.Println(n)
	}

}

func incrementor() chan int {
	one := make(chan int)

	go func() {
		for i := 0; i <10; i++ {
			one <- i
		}
		close(one)
	}()

	return one
}

func puller(one chan int) chan int {
	two := make(chan int)

	go func() {
		var sum int
		for n := range one {
			sum += n
		}
		two <- sum
		close(two)
	}()

	return two
}

/*
the optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.
<- chan means receive From the channel
 */

func incrementorDirection() <- chan int {
	one := make(chan int)

	go func() {
		for i := 0; i <10; i++ {
			one <- i
		}
		close(one)
	}()

	return one
}

func pullerDirection(one <- chan int) <- chan int {
	two := make(chan int)

	go func() {
		var sum int
		for n := range one {
			sum += n
		}
		two <- sum
		close(two)
	}()

	return two
}
