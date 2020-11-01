// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"regexp"
	"strconv"
	"strings"
)

// Add takes two or more numbers and returns the result of adding them together.
func Add(a, b float64, extra ...float64) float64 {
	total := a + b
	for _, n := range extra {
		total += n
	}
	return total
}

// Subtract takes two or more numbers and returns the result of subtracting the
// second and subsequent numbers from the first.
func Subtract(a, b float64, extra ...float64) float64 {
	result := a - b
	for _, n := range extra {
		result -= n
	}
	return result
}

// Multiply takes two or more numbers and returns the result of multiplying them
// together.
func Multiply(a, b float64, extra ...float64) float64 {
	result := a * b
	for _, n := range extra {
		result *= n
	}
	return result
}

// Divide takes two or more numbers and returns the result of dividing the first
// by the second and subsequent numbers, or an error if division by zero occurs.
func Divide(a, b float64, extra []float64) (float64, error) {
	ErrDivideByZero := errors.New("Cannot divide by zero")
	if b == 0 {
		return 0, ErrDivideByZero
	}
	result := a / b
	for _, n := range extra {
		if n == 0 {
			return 0, ErrDivideByZero
		}
		result /= n
	}
	return result, nil
}

// Sqrt takes a number and returns its square root, or an error if the number is negative.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f:  square root of a negative number is not defined", a)
	}
	z := 1.0
	// First guess
	z -= (z*z - a) / (2 * z)
	// Iterate until change is very small
	for zNew, delta := z, z; delta > 0.00000001; z = zNew {
		zNew -= (zNew*zNew - a) / (2 * zNew)
		delta = z - zNew
	}
	return z, nil
}

var validExpression = regexp.MustCompile(`^(\d+)(\.\d+)?(\*|\/|\+|\-)(\d+)(\.\d+)?$`)

// CalculateString takes math formula as string and returns the result in
// float64 format
func CalculateString(input string) (float64, error) {

	input = strings.ReplaceAll(input, " ", "")
	if !validExpression.MatchString(input) {
		return 0, fmt.Errorf("Invalid expression: %q", input)
	}
	input = validExpression.FindString(input)
	fs := token.NewFileSet()
	tr, err := types.Eval(fs, nil, token.NoPos, input)
	if err != nil {
		fmt.Printf("Cannot evaluate expression %s: %e", input, err)
	}
	evaluated, err := strconv.ParseFloat(tr.Value.String(), 64)
	if err != nil {
		fmt.Printf("Cannot convert %q from string to float64: %e", tr.Value.String(), err)
	}

	return evaluated, nil
}
