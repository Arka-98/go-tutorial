package tests

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	res := Add(2, 3)
	exp := 5

	if res != exp {
		t.Errorf("Add(2, 3) = %d, Want %d", res, exp)
	}
}

func TestTableDriven(t *testing.T) {
	tests := []struct{ a, b, exp int }{
		{2, 3, 5},
		{1, 2, 3},
		{8, 2, 9},
	}

	for _, test := range tests {
		res := test.a + test.b

		if res != test.exp {
			t.Errorf("Add(%d, %d) = %d, Want %d", test.a, test.b, res, test.exp)
		}
	}
}

func TestWithSubTests(t *testing.T) {
	tests := []struct{ a, b, exp int }{
		{2, 3, 5},
		{9, 3, 11},
		{3, 5, 8},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			res := test.a + test.b

			if res != test.exp {
				t.Errorf("Result: %d, Want: %d", res, test.exp)
			}
		})
	}
}
