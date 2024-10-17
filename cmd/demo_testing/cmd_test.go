package demo_testing

import "testing"

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
