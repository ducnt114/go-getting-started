package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2) // = num goroutines
	wg.Done()
	pipe := make(chan string)
	go func() {
		for receiver := range pipe {
			fmt.Println(receiver)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			pipe <- fmt.Sprintf("%v", i)
		}
	}()
	close(pipe)
	wg.Wait()
}
