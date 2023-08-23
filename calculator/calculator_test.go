package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 11
	got := calculator.Add(5, 6)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 1
	got := calculator.Subtract(5, 6)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 30
	got := calculator.Multiply(5, 6)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
