package main

import (
	"fmt"
	"math"
)

func countChars(m map[pattern]int) map[rune]int {
	charCount := make(map[rune]int)
	for p, pcount := range m {
		if ccount, ok := charCount[p.l]; !ok {
			charCount[p.l] = pcount
		} else {
			charCount[p.l] = ccount + pcount
		}
		if ccount, ok := charCount[p.r]; !ok {
			charCount[p.r] = pcount
		} else {
			charCount[p.r] = ccount + pcount
		}
	}

	// adjust the numbers
	for r, ccount := range charCount {
		charCount[r] = (ccount + 1) / 2
	}
	return charCount
}

func minmax(m map[rune]int) (min, max int) {
	if len(m) == 0 {
		return
	}
	min, max = math.MaxInt, math.MinInt
	for _, count := range m {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return
}

func nSteps(n int, f formula) int {
	patterns := make(map[pattern]int)
	for i := 1; i < len(f.template); i++ {
		patterns[pattern{f.template[i-1], f.template[i]}] = 1
	}

	for i := 0; i < n; i++ {
		updatedPatterns := make(map[pattern]int, len(patterns))
		for p, pcount := range patterns {
			if r, isRule := f.rules[p]; isRule {
				lpattern := pattern{p.l, r}
				if count, ok := updatedPatterns[lpattern]; !ok {
					updatedPatterns[lpattern] = pcount
				} else {
					updatedPatterns[lpattern] = count + pcount
				}
				rpattern := pattern{r, p.r}
				if count, ok := updatedPatterns[rpattern]; !ok {
					updatedPatterns[rpattern] = pcount
				} else {
					updatedPatterns[rpattern] = count + pcount
				}
			}
		}
		patterns = updatedPatterns
	}

	min, max := minmax(countChars(patterns))
	return max - min
}

func part1(f formula) int {
	return nSteps(10, f)
}

func part2(f formula) int {
	return nSteps(40, f)
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\tinput : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(sample))
	fmt.Printf("\tinput : %d\n", part2(input))

	// Part One
	// 	sample: 1588
	// 	input : 2435
	// Part Two
	// 	sample: 2188189693529
	// 	input : 2587447599164
}
