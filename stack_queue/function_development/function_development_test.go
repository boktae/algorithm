package function_development

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		progresses, speeds, want []int
	}{
		{
			progresses: []int{93, 30, 55},
			speeds:     []int{1, 30, 5},
			want:       []int{2, 1},
		},
		{
			progresses: []int{95, 90, 99, 99, 80, 99},
			speeds:     []int{1, 1, 1, 1, 1, 1},
			want:       []int{1, 3, 2},
		},
	}
	for _, c := range cases {
		got := Solution(c.progresses, c.speeds)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nSolution(%v, %v) == %v, want %v", c.progresses, c.speeds, got, c.want)
		}
	}
}
