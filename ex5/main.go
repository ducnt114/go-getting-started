package main

func main() {
	//t := time.NewTicker(2 * time.Second)
	//defer t.Stop()
	//for {
	//	tick := <-t.C
	//	fmt.Println("tick at: ", tick)
	//}

	// Create a taskChannel with 10 tasks
	tasks := 10
	taskChannel := make(chan int, tasks)

	// Fill the taskChannel with tasks
	for i := 1; i <= tasks; i++ {
		taskChannel <- i
	}
	close(taskChannel)

	// TODO: Create a ticker for rate limiting: 1 task each 1 second
}
