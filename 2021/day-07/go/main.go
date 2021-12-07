package main

import (
	"fmt"
	"math"
	"sort"
)

func cost1(data []int, pos int) int {
	cost := 0
	for _, p := range data {
		cost += int(math.Abs(float64(p) - float64(pos)))
	}
	return cost
}

func part1(data []int) int {
	middle := len(data) / 2
	median := 0
	if len(data)%2 == 1 {
		median = data[int(middle)]
	} else {
		median = int(math.Round((float64(data[middle]) + float64(data[middle-1])) / 2))
	}
	return cost1(data, median)
}

func cost2(data []int, pos int) int {
	cost := 0
	for _, p := range data {
		distance := int(math.Abs(float64(p) - float64(pos)))
		cost += (distance * (distance + 1)) / 2
	}
	return cost
}

func part2(data []int) int {
	min, max := int(data[0]), int(data[len(data)-1])
	cost := cost2(data, min)
	for i := min + 1; i < max; i++ {
		c := cost2(data, i)
		if c > cost {
			break
		}
		cost = c
	}
	return cost
}

func main() {
	sort.Ints(sample)
	sort.Ints(input)

	fmt.Println("Part One")
	fmt.Printf("  sample: %d\n", part1(sample))
	fmt.Printf("  input : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("  sample: %d\n", part2(sample))
	fmt.Printf("  input : %d\n", part2(input))

	// Part One
	//   sample: 37
	//   input : 340987
	// Part Two
	//   sample: 168
	//   input : 96987874
}
