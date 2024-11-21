package main

import "fmt"

func main() {
	n := 10
	ch := make(chan int)

	go fibonacci(n, ch)

	fmt.Println("Fibonacci sequence:")
	for num := range ch {
		fmt.Println(num)
	}
}

func fibonacci(n int, ch chan int) {
	// TODO: implement fibonacci sequence, send the result to the channel
}
