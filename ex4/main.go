package main

import (
	"fmt"
	"time"
)

type Task int
type Result int

type worker struct {
	ID int
}

func (w *worker) processWork(task Task) {
	time.Sleep(1 * time.Second)
	fmt.Printf("workerID: %d, task: %d done\n", w.ID, task)
}

func main() {
	taskChannel := make(chan Task)
	resultChannel := make(chan Result)

	for i := 0; i < 10; i++ {
		t := Task(i)
		taskChannel <- t
	}

	numWorker := 3
	for i := 0; i < numWorker; i++ {
		// TODO: create a worker to process the task
	}

	for r := range resultChannel {
		fmt.Println("result: ", r)
	}
}
