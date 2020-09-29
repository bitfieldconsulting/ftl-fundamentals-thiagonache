// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"go/token"
	"go/types"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, extra ...float64) float64 {
	var total float64 = a + b
	for _, n := range extra {
		total += n
	}
	return total
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, extra []float64) float64 {
	result := a - b
	for _, n := range extra {
		result -= n
	}
	return result
}

// Multiply takes two numbers and returns the result of multiplication them together.
func Multiply(a, b float64, extra []float64) float64 {
	result := a * b
	for _, n := range extra {
		result *= n
	}
	return result
}

// Divide takes two numbers and returns the result of divion and an error message
func Divide(a, b float64, extra []float64) (float64, error) {
	var divideByZeroError = "Cannot divide by zero"
	if b == 0 {
		return 0, errors.New(divideByZeroError)
	}
	result := a / b
	for _, n := range extra {
		if n == 0 {
			return 0, errors.New(divideByZeroError)
		}
		result /= n
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

const regularExpression = `^(\d+)(\.\d+)?(\*|\/|\+|\-)(\d+)(\.\d+)?$`

var regex = regexp.MustCompile(regularExpression)

// CalculateString takes math formula as string and returns the result in float64 format
func CalculateString(input string) (float64, error) {

	noSpace := strings.Replace(input, " ", "", -1)
	match, _ := regexp.Match(regularExpression, []byte(noSpace))
	if match {
		parsed := regex.Find([]byte(noSpace))
		fs := token.NewFileSet()
		tr, err := types.Eval(fs, nil, token.NoPos, string(parsed))
		if err != nil {
			fmt.Printf("Cannot evaluate expression %s: %e", string(parsed), err)
		}
		evaluated, err := strconv.ParseFloat(tr.Value.String(), 64)
		if err != nil {
			fmt.Printf("Cannot convert %q from string to float64: %e", tr.Value.String(), err)
		}
		return evaluated, nil
	}
	return -1, errors.New("Invalid expression")

}
