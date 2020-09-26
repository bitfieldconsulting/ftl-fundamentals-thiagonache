// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"math"
	"strconv"
)

// Add takes two numbers and returns the result of adding them together.
func Add(n1, n2 float64, extra ...float64) float64 {
	var total float64 = n1 + n2
	for i := range extra {
		total += extra[i]
	}
	return total
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(n1, n2 float64, extra []float64) float64 {
	result := n1 - n2
	for i := range extra {
		result -= extra[i]
	}
	return result
}

// Multiply takes two numbers and returns the result of multiplication them together.
func Multiply(n1, n2 float64, extra []float64) float64 {
	result := n1 * n2
	for i := range extra {
		result *= extra[i]
	}
	return result
}

// Divide takes two numbers and returns the result of divion and an error message
func Divide(n1, n2 float64, extra []float64) (float64, error) {
	var divideByZeroError = "Cannot divide by zero"
	if n2 == 0 {
		return 0, errors.New(divideByZeroError)
	}
	result := n1 / n2
	for i := range extra {
		if extra[i] == 0 {
			return 0, errors.New(divideByZeroError)
		}
		result /= extra[i]
	}
	return result, nil
}

// Sqrt takes two numbers and returns the result of divion and an error message
// when receives a negative number
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f:  square root of a negative number is not defined", a)
	}
	return math.Sqrt(a), nil
}

func convertStringFloat64(input string) (float64, error) {
	value, err := strconv.ParseFloat(input, 64)
	return value, err
}

func evaluateExpr(input string) (string, error) {
	fs := token.NewFileSet()
	tr, err := types.Eval(fs, nil, token.NoPos, input)
	if err != nil {
		return "", err
	}
	return tr.Value.String(), err
}

// CalculateString takes math formula as string and returns the result in float64 format
func CalculateString(input string) float64 {
	strValue, err := evaluateExpr(input)
	if err != nil {
		fmt.Printf("Cannot evaluate expression %s: %e", input, err)
	}
	evaluated, err := convertStringFloat64(strValue)
	if err != nil {
		fmt.Printf("Cannot convert %.2f from string to float64: %e", evaluated, err)
	}
	return evaluated
}
