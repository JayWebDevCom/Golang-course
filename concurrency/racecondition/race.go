package racecondition

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int // global variable

// Race is exported
func Race() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		fmt.Println("updating global variable")
		x := counter // increment global variable
		x--          // decrement global variable
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
