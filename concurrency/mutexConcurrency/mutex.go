package mutexConcurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
mutex eliminates race condition because protects access to the mutexCountervariable
*/

var wg sync.WaitGroup
var mutex sync.Mutex
var mutexCounter int

// Mutex is exported
func Mutex() {
	wg.Add(2)
	go incrementor("Foo mutex:")
	go incrementor("Bar mutex:")
	wg.Wait()
	fmt.Println("Final Counter:", mutexCounter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()

		mutexCounter++
		fmt.Println(s, i, "mutexCounter:", mutexCounter)
		mutex.Unlock()
	}
	wg.Done()
}
