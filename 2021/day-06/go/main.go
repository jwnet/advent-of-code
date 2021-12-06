package main

import "fmt"

func fishBreeding(fish []int, ndays int) int {
	fishCycle := [9]int{}
	for _, fish := range fish {
		fishCycle[fish]++
	}
	for i := 0; i < ndays; i++ {
		updatedCycle := [9]int{}
		for day, nfish := range fishCycle {
			if day == 0 {
				updatedCycle[6] += nfish
				updatedCycle[8] = nfish
			} else {
				updatedCycle[day-1] += nfish
			}
		}
		fishCycle = updatedCycle
	}
	total := 0
	for _, count := range fishCycle {
		total += count
	}
	return total
}

func part1(data []int) int {
	return fishBreeding(data, 80)
}

func part2(data []int) int {
	return fishBreeding(data, 256)
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("  sample: %d\n", part1(sample)) // 5934
	fmt.Printf("  input : %d\n", part1(input))  // 377263
	fmt.Println("Part Two")
	fmt.Printf("  sample: %d\n", part2(sample)) // 26984457539
	fmt.Printf("  input : %d\n", part2(input))  // 1695929023803
}
