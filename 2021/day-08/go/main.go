package main

import (
	"fmt"
	"strings"
)

func decodePatterns(patterns []string) map[string]string {
	decoder := make(map[string]string, 10)
	helper := make(map[int]string, 4)
	remaining := make([]string, 0, 6)
	for _, p := range patterns {
		switch len(p) {
		case 2:
			decoder[p] = "1"
			helper[1] = p
		case 3:
			decoder[p] = "7"
			helper[7] = p
		case 4:
			decoder[p] = "4"
			helper[4] = p
		case 7:
			decoder[p] = "8"
			helper[8] = p
		default:
			remaining = append(remaining, p)
		}
	}
	for _, p := range remaining {
		switch len(p) {
		case 5: // 2, 3, 5
			switch {
			case strings.ContainsRune(p, helper[1][0]) && strings.ContainsRune(p, helper[1][1]):
				decoder[p] = "3"
			case 
			}
		case 6: // 0, 6, 9
		}
	}
	return decoder
}

func part1(data []signal) int {
	count1478 := 0
	for _, s := range data {
		for _, n := range s.output {
			switch len(n) {
			case 2, 3, 4, 7:
				count1478++
			}
		}
	}
	return count1478
}

func part2(data []signal) int {

	return -1
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("  sample: %d\n", part1(sample))
	fmt.Printf("  input : %d\n", part1(input))
	// fmt.Println("Part Two")
	// fmt.Printf("  sample: %d\n", part2(sample))
	// fmt.Printf("  input : %d\n", part2(input))
}
