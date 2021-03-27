package lagest_number

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		numbers []int
		result  string
	}{
		{
			numbers: []int{6, 10, 2},
			result:  "6210",
		},
		{
			numbers: []int{3, 30, 34, 5, 9},
			result:  "9534330",
		},
		{
			numbers: []int{0, 0, 0, 0, 0},
			result:  "0",
		},
	}
	for _, c := range cases {
		got := Solution(c.numbers)

		if got != c.result {
			t.Errorf("\nSolution(%v) == %v, want %v", c.numbers, got, c.result)
		}
	}
}
