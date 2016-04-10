package main

import "errors"

var ErrDivisionByZero = errors.New("Division by zero.")

// AddNumbers takes 2 numbers and returns the sum.
func AddNumbers(a, b int) int {
	return a + b
}

// AddNumbers takes 2 numbers and returns the difference.
func RestNumbers(a, b int) int {
	return a - b
}

// DivNumbers takes 2 numbers and performs integer division.
func DivNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
