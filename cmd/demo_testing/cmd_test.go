package demo_testing

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// BenchmarkConcatWithPlus benchmarks the ConcatWithPlus function
//func BenchmarkConcatWithPlus(b *testing.B) {
//	strs := []string{"Go", "is", "awesome", "and", "fast!"}
//	for i := 0; i < b.N; i++ {
//		ConcatWithPlus(strs)
//	}
//}

// BenchmarkConcatWithBuilder benchmarks the ConcatWithBuilder function
//func BenchmarkConcatWithBuilder(b *testing.B) {
//	strs := []string{"Go", "is", "awesome", "and", "fast!"}
//	for i := 0; i < b.N; i++ {
//		ConcatWithBuilder(strs)
//	}
//}

type InputJson struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func BenchmarkJsonUnmarshall(b *testing.B) {
	input := []byte(`{"id":1,"name":"John","age":30,"password":"secret"}`)
	for i := 0; i < b.N; i++ {
		JsonUnmarshall(input, &InputJson{})
	}
}

func BenchmarkJsonUnmarshallWithLib(b *testing.B) {
	input := []byte(`{"id":1,"name":"John","age":30,"password":"secret"}`)
	for i := 0; i < b.N; i++ {
		JsonUnmarshallWithLib(input, &InputJson{})
	}
}

func TestCtxTimeout(t *testing.T) {
	// Set a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Important: release resources if operation completes before timeout

	// Run a function with a simulated delay
	err := performTask(ctx)
	if err != nil {
		fmt.Println("tas Error: ", err)
	} else {
		fmt.Println("Task completed successfully.")
	}
}

func performTask(ctx context.Context) error {
	// Simulating a task that takes 3 seconds
	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("task exceeded the allowed time limit")
	case <-ctx.Done():
		// The context's Done channel will be closed when the timeout expires
		return ctx.Err()
	}
}
