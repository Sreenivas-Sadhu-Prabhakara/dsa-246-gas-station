package main

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	cases := []struct {
		gas, cost []int
		want      int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{5}, []int{4}, 0},
		{[]int{5}, []int{6}, -1},
		{[]int{3, 1, 1}, []int{1, 2, 2}, 0},
		{[]int{0, 0, 5}, []int{1, 1, 1}, 2},
	}
	for _, c := range cases {
		if got := CanCompleteCircuit(c.gas, c.cost); got != c.want {
			t.Errorf("CanCompleteCircuit(%v,%v) = %d, want %d", c.gas, c.cost, got, c.want)
		}
	}
}
