package main

import (
	"fmt"
	"math"
)

func part1(data []float64) int {
	//find median then compute cost
}

func part2(data) int {
	// sort first!!!
	x := data[0]
	for i := 1; i < len(data); i++ {
		x = (data[i] + x) / 2
	}
	pos = int(math.Round(x))
	//compute cost
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("  sample: %d\n", part1(sample))
	// fmt.Printf("  input : %d\n", part1(input))
	// fmt.Println("Part Two")
	// fmt.Printf("  sample: %d\n", part2(sample))
	// fmt.Printf("  input : %d\n", part2(input))
}
