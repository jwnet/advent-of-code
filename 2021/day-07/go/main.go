package main

import (
	"fmt"
	"math"
	"sort"
)

func part1(data []float64) int {
	sort.Float64s(data)
	middle := float64(len(data)) / 2
	median := 0.0
	if len(data) % 2 == 1 {
		median = data[int(middle)]
	} else {
		median = math.Round((data[int(middle) - 1] + data[int(middle)]) / 2)
	}

	cost := 0
	for _, pos := range data {
		cost += int(math.Abs(pos - median))
	}
	return cost
}

// func part2(data) int {

// }

func main() {
	fmt.Println("Part One")
	fmt.Printf("  sample: %d\n", part1(sample))
	fmt.Printf("  input : %d\n", part1(input))
	// fmt.Println("Part Two")
	// fmt.Printf("  sample: %d\n", part2(sample))
	// fmt.Printf("  input : %d\n", part2(input))
}
