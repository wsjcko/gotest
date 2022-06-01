package main

import (
	"math/rand"
	"testing"
)

func generateSlice() []int64 {
	s := make([]int64, 0, 1)
	for i := 0; i < 1000000; i++ {
		s = append(s, rand.Int63())
	}
	return s
}

func BenchmarkRemoveDuplicatesInPlace(b *testing.B) {
	s := generateSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveDuplicatesInPlace(s)
	}
}
