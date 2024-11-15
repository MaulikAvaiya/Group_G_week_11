package main

import "testing"

// Benchmark for Divide function with non-zero denominators
func BenchmarkDivision(b *testing.B) {
	for h := 0; h < b.N; h++ {
		Division(30, 7)
	}
}

// Benchmark Square with representative inputs.
func BenchmarkSquare(b *testing.B) {
	for h := 0; h < b.N; h++ {
		Square(29)
	}
}
