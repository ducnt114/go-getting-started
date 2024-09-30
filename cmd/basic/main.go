package main

import "fmt"

func main() {
	//LearDefer()
	//LearnSlice()

	LearRecover()
}

// LearDefer
// defer ~ stack: last in - first out
func LearDefer() {
	fmt.Println("command 1")
	defer func() {
		fmt.Println("command 2")
	}()

	fmt.Println("command 3")

	defer fmt.Println("command 4")

	fmt.Println("command 5")
}

// LearnSlice
// len and capacity
func LearnSlice() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(a)
	fmt.Println(b)
}

func LearRecover() {
	fmt.Println(calc(4, 2))
	fmt.Println(calc(4, 0))
	fmt.Println(calc(6, 2))
}

func calc(a, b int) (res int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			res = -1
		}
	}()

	res = a / b
	return res
}
