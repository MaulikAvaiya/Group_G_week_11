package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := Divide(6, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide result:", result)
	}
	expected := 1
	if result != expected {
		t.Errorf("Divide(6, 0) will result %d, but we got %d", result, expected)
	}
}

func TestSquare(t *testing.T) {
	result := Square(2)
	expected := 4
	if result != expected {
		t.Errorf("Square(2) will result %d, but we got %d", result, expected)
	}
}

// Benchmark for Divide function with non-zero denominators
func BenchmarkDivide(b *testing.B) {
	for h := 0; h < b.N; h++ {
		Divide(30, 7)
	}
}

// Benchmark Square with representative inputs.
func BenchmarkSquare(b *testing.B) {
	for h := 0; h < b.N; h++ {
		Square(29)
	}
}
