package main

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
