package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

// increment increments the counter with a mutex lock
func increment(wg *sync.WaitGroup) {
	// Locking the mutex to ensure only one goroutine accesses the counter
	//mu.Lock()
	counter++
	fmt.Printf("Counter: %d\n", counter)
	// Unlocking the mutex to allow other goroutines to access the counter
	//mu.Unlock()
	// Signal that this goroutine is done
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	// Starting 5 goroutines to increment the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Waiting for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter Value:", counter)
}
