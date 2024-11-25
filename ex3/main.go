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

// 1,2,3,5,8.....
func fibonacci(n int, ch chan int) {
	a := 1
	b := 2
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}
