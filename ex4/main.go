package main

import (
	"fmt"
)

type Task int
type Result int

type worker struct {
	ID         int
	ch         chan Task
	resultChan chan Result
}

func (w *worker) processWork() {
	t := <-w.ch
	w.resultChan <- Result(t)
	fmt.Printf("workerID: %d, task: %d done\n",
		w.ID, t)
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
