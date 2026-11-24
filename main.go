package main

import "fmt"

// CanCompleteCircuit returns the unique start index from which the circular
// route can be driven (tank += gas[i] - cost[i] at each station), or -1.
// One pass: a negative running tank disqualifies every start in the current
// stretch, so the candidate restarts just past the failure.
func CanCompleteCircuit(gas, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff
		if tank < 0 {
			start = i + 1 // no station in [start..i] can be the answer
			tank = 0
		}
	}
	if total < 0 {
		return -1 // overall deficit: no start works
	}
	return start
}

func main() {
	fmt.Println(CanCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(CanCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
