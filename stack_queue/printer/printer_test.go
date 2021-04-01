package printer

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		priorities []int
		location   int
		result     int
	}{
		{
			priorities: []int{2, 1, 3, 2},
			location:   2,
			result:     1,
		},
		{
			priorities: []int{1, 1, 9, 1, 1, 1},
			location:   0,
			result:     5,
		},
		{
			priorities: []int{6, 1, 1, 1, 1, 1},
			location:   0,
			result:     1,
		},
		{
			priorities: []int{6, 1, 1, 1, 1, 1},
			location:   1,
			result:     2,
		},
		{
			priorities: []int{1, 1, 6},
			location:   2,
			result:     1,
		},
		{
			priorities: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			location:   7,
			result:     2,
		},
		{
			priorities: []int{2, 1, 3, 1},
			location:   2,
			result:     1,
		},
	}
	for _, c := range cases {
		got := Solution(c.priorities, c.location)

		if got != c.result {
			t.Errorf("\nSolution(%v, %v) == %v, want %v", c.priorities, c.location, got, c.result)
		}
	}
}
