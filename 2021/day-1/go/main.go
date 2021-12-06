package main

import (
	"fmt"
)

// part one
func countIncreases(data []int) int {
	count := 0
	for i, num := range data {
		if i == 0 {
			continue
		}
		if num > input[i-1] {
			count++
		}
	}
	return count
}

// part two
func countGroupIncreases(data []int) int {
	count := 0
	for i := 3; i < len(data); i++ {
		a := input[i-3] + input[i-2] + input[i-1]
		b := input[i-2] + input[i-1] + input[i]
		if b > a {
			count++
		}
	}
	return count
}

func main() {
	// part one
	fmt.Println(countIncreases(input)) // output: 1688

	// part two
	fmt.Println(countGroupIncreases(input)) // output: 1728
}
