package atomicity

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
small locked thing
go run -race main.go - shouldn't get race condition warning but do ?
*/

var wg sync.WaitGroup
var atomicityCounter int64

var cf = func(add int64) {
	atomic.AddInt64(&atomicityCounter, add)
}

// Atomicity is exported
func Atomicity() {
	wg.Add(2)
	go incrementor("Foo atomic:", cf)
	go incrementor("Bar atomic:", cf)
	wg.Wait()
	fmt.Println("Final Counter:", atomicityCounter)
}

func incrementor(s string, someConcurrencyFunction func(add int64)) {
	for i := 0; i < 20; i++ {

		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		someConcurrencyFunction(int64(i))

		fmt.Println(s, s, i, s, "atomiticyCounter:", atomicityCounter)
	}
	wg.Done()
}
