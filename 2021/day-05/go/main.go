package main

import "fmt"

func part1(lines []line) int {
	var grid [1000][1000]int
	intersections := 0
	for _, l := range lines {
		if l.diagonal() {
			continue
		}
		for _, p := range l.points() {
			grid[p.y][p.x]++
			if grid[p.y][p.x] == 2 {
				intersections++
			}
		}
	}
	return intersections
}

func part2(lines []line) int {
	var grid [1000][1000]int
	intersections := 0
	for _, l := range lines {
		for _, p := range l.points() {
			grid[p.y][p.x]++
			if grid[p.y][p.x] == 2 {
				intersections++
			}
		}
	}
	return intersections
}

func main() {
	fmt.Println("Part One")
	fmt.Println("  sample: ", part1(sample)) // 5
	fmt.Println("  input : ", part1(input))  // 6856

	fmt.Println("Part Two")
	fmt.Println("  sample:", part2(sample)) // 12
	fmt.Println("  input : ", part2(input)) // 20666
}
