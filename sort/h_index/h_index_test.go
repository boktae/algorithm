package h_index

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		citations []int
		result    int
	}{
		{
			citations: []int{3, 0, 6, 1, 5},
			result:    3,
		},
		{
			citations: []int{10, 8, 5, 4, 3},
			result:    4,
		},
		{
			citations: []int{25, 8, 5, 3, 3},
			result:    3,
		},
	}
	for _, c := range cases {
		got := Solution(c.citations)

		if got != c.result {
			t.Errorf("\nSolution(%v) == %v, want %v", c.citations, got, c.result)
		}
	}
}
