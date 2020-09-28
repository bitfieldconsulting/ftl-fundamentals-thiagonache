package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("Fixed number addition", func(t *testing.T) {
		var want float64 = 4
		got := calculator.Add(2, 2)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	})

	t.Run("Sum four numbers", func(t *testing.T) {
		var want float64 = 40
		got := calculator.Add(2, 8, 20, 10)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	})

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
		n1    float64
		n2    float64
		extra []float64
	}{
		{
			name:  "Substract a > b returns positive",
			n1:    3,
			n2:    2,
			extra: nil,
			want:  1,
		},
		{
			name:  "Substract a < b returns negative",
			n1:    3,
			n2:    5,
			extra: nil,
			want:  -2,
		},
		{
			name:  "Ensure decimals",
			n1:    6,
			n2:    2.225,
			extra: nil,
			want:  3.77500,
		},
		{
			name:  "Substract a and b negatives which returns a positive",
			n1:    -3,
			n2:    -5,
			extra: nil,
			want:  2,
		},
		{
			name:  "Substract five numbers",
			n1:    20,
			n2:    5,
			extra: []float64{4, 3, 7},
			want:  1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Subtract(tC.n1, tC.n2, tC.extra)
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
		n1    float64
		n2    float64
		extra []float64
	}{
		{
			name:  "Multiply two positive numbers returns positive",
			n1:    3,
			n2:    2,
			extra: nil,
			want:  6,
		},
		{
			name:  "Multiply positive and negative should return negative",
			n1:    3,
			n2:    -2,
			extra: nil,
			want:  -6,
		},
		{
			name:  "Multiply two negative numbers returns positive",
			n1:    -3,
			n2:    -2,
			extra: nil,
			want:  6,
		},
		{
			name:  "Multiply fractions return fraction",
			n1:    3.2,
			n2:    2.501,
			extra: nil,
			want:  8.0032,
		},
		{
			name:  "Multiply by zero return zero",
			n1:    1943,
			n2:    0,
			extra: nil,
			want:  0,
		},
		{
			name:  "Multiply five numbers",
			n1:    2,
			n2:    2,
			extra: []float64{5, 3, 2},
			want:  120,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.Multiply(tC.n1, tC.n2, tC.extra)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

// It only consider errExpected variable if an error occours.
func TestDivide(t *testing.T) {
	testCases := []struct {
		name        string
		errExpected bool
		want        float64
		n1          float64
		n2          float64
		extra       []float64
	}{
		{
			name:        "Divide two positive integers",
			errExpected: false,
			want:        2,
			n1:          6,
			n2:          3,
			extra:       nil,
		},
		{
			name:        "Divide by zero",
			errExpected: true,
			want:        0,
			n1:          6,
			n2:          0,
			extra:       nil,
		},
		{
			name:        "Divide two negative integers",
			errExpected: false,
			want:        3,
			n1:          -6,
			n2:          -2,
			extra:       nil,
		},
		{
			name:        "Divide five numbers",
			errExpected: false,
			want:        1,
			n1:          60,
			n2:          2,
			extra:       []float64{3, 5, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := calculator.Divide(tC.n1, tC.n2, tC.extra)
			if err != nil && tC.errExpected == false {
				t.Fatalf("Cannot divide inputs %.2f %.2f %.2f: %s", tC.n1, tC.n2, tC.extra, err)
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
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestCalculateString(t *testing.T) {
	testCases := []struct {
		name    string
		want    float64
		formula string
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
			name:    "Substract fraction with no space ",
			want:    99.9,
			formula: "100-0.1",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := calculator.CalculateString(tC.formula)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}
