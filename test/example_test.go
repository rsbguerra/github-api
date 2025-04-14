package test

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, expected)
	}
}
