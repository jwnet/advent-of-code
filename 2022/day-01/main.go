package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// part 1: 70764
// part 2: 203905
func main() {
	max := 0
	for _, list := range input {
		total := sum(list)
		if total > max {
			max = total
		}
	}
	fmt.Printf("part 1: %d\n", max)

	top3 := []int{0, 0, 0}
	for _, list := range input {
		total := sum(list)
		if yes, idx := inTop3(total, top3); yes {
			insertShift(top3, total, idx)
		}
	}
	fmt.Printf("part 2: %d\n", sum(top3))
}

func inTop3(cals int, top3 []int) (yes bool, idx int) {
	for i, val := range top3 {
		if cals >= val {
			idx = i
			yes = true
			return
		}
	}
	return
}

func sum[T constraints.Integer](s []T) T {
	var total T
	for _, val := range s {
		total += val
	}
	return total
}

func insertShift[T any](s []T, val T, idx int) {
	prev := s[idx]
	s[idx] = val
	for i := idx + 1; i < len(s); i++ {
		prev, s[i] = s[i], prev
	}
}
