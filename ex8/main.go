package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)
}

func PrintFn(x int) {
	fmt.Println(x)
}
