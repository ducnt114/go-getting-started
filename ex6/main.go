package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for item := range chan1 {

		}
		for {
			item := <-chan1
		}
		for {
			select {
			case item := <-chan1:
				fmt.Println(item)
			case item := <-chan2:
				fmt.Println(item)
			}
		}
		switch expr {
		case 1, 2, 3, 4:
			
		case 4, 5, 6:
		default:
		}

		for i := 0; i < 10; i++ {
			chan1 <- i
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
		}
	}()
	for i := 0; i < 20; i++ {
		// TODO use select to receive from chan1 and chan2
	}
	close(chan1)
	close(chan2)
}
