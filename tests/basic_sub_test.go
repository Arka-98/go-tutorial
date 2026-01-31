package tests

import "testing"

func Subtract(a, b int) int {
	return a - b
}

func TestSubtract(t *testing.T) {
	res := Subtract(5, 2)
	exp := 3

	if res != exp {
		t.Errorf("Subtract(5, 2) = %d, Want %d", res, exp)
	}
}
