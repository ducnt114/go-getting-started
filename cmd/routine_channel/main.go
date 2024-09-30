package main

import (
	"fmt"
	"sync"
)

func main() {
	//LearnChannel()
	LearnWaitGroup()
}

func LearnChannel() {
	var c chan int
	c = make(chan int)

	go worker(c)
	go worker(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}
	//time.Sleep(5 * time.Second)
	//close(c)
}

func worker(c chan int) {
	for x := range c {
		fmt.Println(x)
	}
}

func LearnWaitGroup() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printFunc(i, wg)
	}
	wg.Wait()
}

func printFunc(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
