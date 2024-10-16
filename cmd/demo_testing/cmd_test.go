package demo_testing

import "testing"

// BenchmarkConcatWithPlus benchmarks the ConcatWithPlus function
func BenchmarkConcatWithPlus(b *testing.B) {
	strs := []string{"Go", "is", "awesome", "and", "fast!"}
	for i := 0; i < b.N; i++ {
		ConcatWithPlus(strs)
	}
}

// BenchmarkConcatWithBuilder benchmarks the ConcatWithBuilder function
func BenchmarkConcatWithBuilder(b *testing.B) {
	strs := []string{"Go", "is", "awesome", "and", "fast!"}
	for i := 0; i < b.N; i++ {
		ConcatWithBuilder(strs)
	}
}
