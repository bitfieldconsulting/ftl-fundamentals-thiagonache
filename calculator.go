// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"math"
)

var divideByZeroMessage = "Cannot divide by zero"
var negativeNumber = "No negative number allowed"

// ErrorString is a trivial implementation of error.
type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func extractFirstItem(numbers []float64) (float64, []float64) {
	return numbers[0], numbers[1:]
}

// Add takes two numbers and returns the result of adding them together.
func Add(numbers ...float64) float64 {
	var total float64 = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(numbers []float64) float64 {
	firstValue, values := extractFirstItem(numbers)
	for i := 0; i < len(values); i++ {
		firstValue -= values[i]
	}
	return firstValue
}

// Multiply takes two numbers and returns the result of multiplication them together.
func Multiply(numbers []float64) float64 {
	firstValue, values := extractFirstItem(numbers)
	for i := 0; i < len(values); i++ {
		firstValue *= values[i]
	}
	return firstValue
}

// Divide takes two numbers and returns the result of divion and an error message
func Divide(numbers []float64) (float64, error) {
	firstValue, values := extractFirstItem(numbers)
	for i := 0; i < len(values); i++ {
		if values[i] == 0 {
			return 0, &ErrorString{divideByZeroMessage}
		}
		firstValue /= values[i]
	}
	return firstValue, nil
}

// Sqrt takes two numbers and returns the result of divion and an error message
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, &ErrorString{negativeNumber}
	}
	return math.Sqrt(a), nil
}
