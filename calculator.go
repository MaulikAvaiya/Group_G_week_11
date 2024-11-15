package main

import (
	"errors"
	"fmt"
)

// Divide takes two integers and returns their quotient and an error if division by zero is attempted.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Square takes an integer and returns its square.
func Square(a int) int {
	return a * a
}

func main() {
	// Test the Divide function
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide result:", result)
	}

	// Test the Square function
	squareResult := Square(5)
	fmt.Println("Square result:", squareResult)
}
