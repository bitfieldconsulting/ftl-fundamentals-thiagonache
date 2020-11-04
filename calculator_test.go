package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name  string
		want  float64
		a     float64
		b     float64
		extra []float64
	}{
		{
			name: "Add two numbers",
			want: 4,
			a:    2,
			b:    2,
		},
		{
			name:  "Sum four numbers",
			want:  40,
			a:     2,
			b:     8,
			extra: []float64{20, 10},
		},
		{
			name: "Sum negative numbers",
			want: -32,
			a:    -16,
			b:    -16,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Add(tC.a, tC.b, tC.extra...)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	var totalRandom int = 100
	t.Run("Random number addition", func(t *testing.T) {
		for count := 0; count < totalRandom; count++ {
			var a float64 = rand.NormFloat64()*1024 + 2
			var b float64 = rand.NormFloat64()*2048 + 5
			want := a + b
			got := calculator.Add(a, b)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		}
	})
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		name  string
		want  float64
		a     float64
		b     float64
		extra []float64
	}{
		{
			name: "Subtract a > b returns positive",
			a:    3,
			b:    2,
			want: 1,
		},
		{
			name: "Subtract a < b returns negative",
			a:    3,
			b:    5,
			want: -2,
		},
		{
			name: "Ensure decimals",
			a:    6,
			b:    2.225,
			want: 3.77500,
		},
		{
			name: "Subtract a and b negatives which returns a positive",
			a:    -3,
			b:    -5,
			want: 2,
		},
		{
			name:  "Subtract five numbers",
			a:     20,
			b:     5,
			extra: []float64{4, 3, 7},
			want:  1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Subtract(tC.a, tC.b, tC.extra...)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name  string
		want  float64
		a     float64
		b     float64
		extra []float64
	}{
		{
			name: "Multiply two positive numbers returns positive",
			a:    3,
			b:    2,
			want: 6,
		},
		{
			name: "Multiply positive and negative should return negative",
			a:    3,
			b:    -2,
			want: -6,
		},
		{
			name: "Multiply two negative numbers returns positive",
			a:    -3,
			b:    -2,
			want: 6,
		},
		{
			name: "Multiply fractions return fraction",
			a:    3.2,
			b:    2.501,
			want: 8.0032,
		},
		{
			name: "Multiply by zero return zero",
			a:    1943,
			b:    0,
			want: 0,
		},
		{
			name:  "Multiply five numbers",
			a:     2,
			b:     2,
			extra: []float64{5, 3, 2},
			want:  120,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Multiply(tC.a, tC.b, tC.extra...)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		want        float64
		a           float64
		b           float64
		extra       []float64
	}{
		{
			name: "Divide two positive integers",
			want: 2,
			a:    6,
			b:    3,
		},
		{
			name:        "Divide by zero",
			errExpected: true,
			want:        0,
			a:           6,
			b:           0,
		},
		{
			name: "Divide two negative integers",
			want: 3,
			a:    -6,
			b:    -2,
		},
		{
			name:  "Divide five numbers",
			want:  1,
			a:     60,
			b:     2,
			extra: []float64{3, 5, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b, tC.extra)
			errReceived := err != nil
			if tC.errExpected != errReceived {
				t.Fatalf("Divide(%f, %f, %f): unexpected error status: %v", tC.a, tC.b, tC.extra, err)
			}

			if tC.want != got {
				t.Errorf("Want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		want, input float64
	}{
		{
			name:  "Calculate square root of a positive number",
			input: 49,
			want:  7,
		},
		{
			name:  "Calculate square root of a long positive number",
			input: 94339,
			want:  307.1465448283604,
		},
		{
			name:        "Calculate square root of a negative number",
			errExpected: true,
			input:       -49,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tC.input)
			errReceived := err != nil
			if tC.errExpected != errReceived {
				t.Fatalf("Sqrt(%.2f): unexpected error status %v", tC.input, err)
			}
			// Check if difference between two numbers is smaller than precision
			opt := cmp.Comparer(func(x, y float64) bool {
				delta := math.Abs(x - y)
				mean := math.Abs(x+y) / 2.0
				return delta/mean < calculator.Precision
			})
			if !errReceived && !cmp.Equal(tC.want, got, opt) {
				t.Error(cmp.Diff(tC.want, got))
			}
		})
	}
}

func TestCalculateString(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		want        float64
		formula     string
	}{
		{
			name:    "Multiply no space",
			want:    4.0,
			formula: "2*2",
		},
		{
			name:    "Sum with fraction",
			want:    2.5,
			formula: "1 + 1.5",
		},
		{
			name:    "Divide with spaces",
			want:    3,
			formula: "18   /   6",
		},
		{
			name:    "Subtract fraction with no space ",
			want:    99.9,
			formula: "100-0.1",
		},
		{
			name:    "Subtract fraction with no space ",
			want:    99.9,
			formula: "100-0.1",
		},
		{
			name:        "Invalid expression",
			errExpected: true,
			want:        1,
			formula:     "1 * 1 * 1",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := calculator.CalculateString(tC.formula)
			errReceived := err != nil
			if tC.errExpected != errReceived {
				t.Fatalf("CaculateString(%s): unexpected error status %v", tC.formula, err)
			}
			if !tC.errExpected && tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}
