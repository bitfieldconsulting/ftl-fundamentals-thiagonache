package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func assertNumbers(want, got float64, t *testing.T) {
	if want != got {
		t.Fatalf("want %f, got %f", want, got)
	}
}

func assertNumbersErrHandling(want, got float64, errExpected bool, t *testing.T) {
	if (want != got) && !errExpected {
		t.Fatalf("want %f, got %f", want, got)
	}
}

func TestAdd(t *testing.T) {
	var totalRandom int = 100

	t.Run("Fixed number addition", func(t *testing.T) {
		var want float64 = 4
		got := calculator.Add(2, 2)
		assertNumbers(want, got, t)
	})

	t.Run("Sum four numbers", func(t *testing.T) {
		var want float64 = 40
		got := calculator.Add(2, 8, 20, 10)
		assertNumbers(want, got, t)
	})

	t.Run("Random number addition", func(t *testing.T) {
		for count := 0; count < totalRandom; count++ {
			var a float64 = rand.NormFloat64()*1024 + 2
			var b float64 = rand.NormFloat64()*2048 + 5
			want := a + b
			got := calculator.Add(a, b)
			assertNumbers(want, got, t)
		}
	})

}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		name   string
		want   float64
		inputs []float64
	}{
		{
			name:   "Substract a > b returns positive",
			inputs: []float64{3, 2},
			want:   1,
		},
		{
			name:   "Substract a < b returns negative",
			inputs: []float64{3, 5},
			want:   -2,
		},
		{
			name:   "Ensure decimals",
			inputs: []float64{6, 2.225},
			want:   3.77500,
		},
		{
			name:   "Substract a and b negatives which returns a positive",
			inputs: []float64{-3, -5},
			want:   2,
		},
		{
			name:   "Substract five numbers",
			inputs: []float64{10, 5, 4},
			want:   1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Subtract(tC.inputs)
			assertNumbers(tC.want, got, t)
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name   string
		want   float64
		inputs []float64
	}{
		{
			name:   "Multiply two positive numbers returns positive",
			inputs: []float64{3, 2},
			want:   6,
		},
		{
			name:   "Multiply positive and negative should return negative",
			inputs: []float64{3, -2},
			want:   -6,
		},
		{
			name:   "Multiply two negative numbers returns positive",
			inputs: []float64{-3, -2},
			want:   6,
		},
		{
			name:   "Multiply fractions return fraction",
			inputs: []float64{3.2, 2.501},
			want:   8.0032,
		},
		{
			name:   "Multiply by zero return zero",
			inputs: []float64{1943, 0},
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Multiply(tC.inputs)
			assertNumbers(tC.want, got, t)
		})
	}
}

// It only consider errExpected variable if an error occours.
func TestDivide(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		want        float64
		inputs      []float64
	}{
		{
			name:        "Divide two positive integers",
			errExpected: false,
			want:        2,
			inputs:      []float64{6, 3},
		},
		{
			name:        "Divide by zero",
			errExpected: true,
			want:        0,
			inputs:      []float64{6, 0},
		},
		{
			name:        "Divide two negative integers",
			errExpected: true,
			want:        -6,
			inputs:      []float64{-2, -3},
		},
		{
			name:        "Divide five numbers",
			errExpected: false,
			want:        1,
			inputs:      []float64{60, 2, 3, 5, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := calculator.Divide(tC.inputs)
			if err != nil && tC.errExpected == false {
				t.Fatalf("Cannot divide inputs %.2f: %s", tC.inputs, err)
			}
			assertNumbersErrHandling(tC.want, got, tC.errExpected, t)
		})
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		want, a     float64
	}{
		{
			name:        "Calculate square root of a positive number",
			errExpected: false,
			want:        7,
			a:           49,
		},
		{
			name:        "Calculate square root of a negative number",
			errExpected: true,
			want:        0,
			a:           -49,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tC.a)
			if err != nil && tC.errExpected == false {
				t.Fatalf("Cannot calculate square root of %f: %s", tC.a, err)
			}
			assertNumbersErrHandling(tC.want, got, tC.errExpected, t)
		})
	}
}
