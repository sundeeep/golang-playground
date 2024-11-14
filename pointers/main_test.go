package main

// Command to run the benchmarks by turning off go compiler optimization:
// go test -bench=. -count 1 -gcflags=-N

import (
	"testing"
)

// Benchmark of Pass by Pointer
func BenchmarkPBP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		funcPBP(&obj)
	}
}

// Benchmark of Pass by Value
func BenchmarkPBV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		funcPBV(obj)
	}
}
