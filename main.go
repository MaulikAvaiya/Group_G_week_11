package main

import (
	"fmt"
)

func Division(x, y float32) (float32, error) {
	if y == 0 {
		return 0, fmt.Errorf("It cannot be divided by zero")
	}
	return x / y, nil
}
func Square(x float32) (float32, error) {
	return x * x, nil

}
