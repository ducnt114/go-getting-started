package service

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCtxTimeout(t *testing.T) {
	s := &bookServiceImpl{}
	if err := s.GetData(context.Background()); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSelectTimeout(t *testing.T) {
	counter := 0
	isTimeOut := false
	for {
		select {
		case <-time.After(1 * time.Second):
			t.Log("timeout after 1s")
			isTimeOut = true
		case <-time.After(2 * time.Second):
			t.Log("timeout after 2s")
			isTimeOut = true
			//default:
			//	counter++
		}
		counter++
		if isTimeOut {
			break
		}
	}
	t.Logf("counter: %v", counter)
}

func f(left, right chan int) {
	left <- 1 + <-right
}

func TestDaisyChain(t *testing.T) {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
