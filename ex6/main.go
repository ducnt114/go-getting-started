package main

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
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
