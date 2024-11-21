package main

import "fmt"

func main() {
	pipe := make(chan string)
	go func() {
		for receiver := range pipe {
			fmt.Println(receiver)
		}
	}()
	pipe <- "water 1"
	pipe <- "water 2"
	pipe <- "water 3"
	close(pipe)
	// TODO: wait for the goroutine to finish
}
