package truck_passing_the_bridge

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		bridge_length, weight, want int
		truck_weights               []int
	}{
		{
			bridge_length: 2,
			weight:        10,
			truck_weights: []int{7, 4, 5, 6},
			want:          8,
		},
		{
			bridge_length: 100,
			weight:        100,
			truck_weights: []int{10},
			want:          101,
		},
		{
			bridge_length: 100,
			weight:        100,
			truck_weights: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			want:          110,
		},
	}
	for _, c := range cases {
		got := Solution(c.bridge_length, c.weight, c.truck_weights)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nSolution(%v, %v, %v) == %v, want %v", c.bridge_length, c.weight, c.truck_weights, got, c.want)
		}
	}
}
