package main

import (
	"fmt"
	"sort"
)

type point struct {
	row, col int
}

func lowPoints(heightmap [][]int) []point {
	lp := []point{}
	for i, row := range heightmap {
		for j, height := range row {
			//left
			if j > 0 && height >= row[j-1] {
				continue
			}
			// right
			if j < len(row)-1 && height >= row[j+1] {
				continue
			}
			// up
			if i > 0 && height >= heightmap[i-1][j] {
				continue
			}
			// down
			if i < len(heightmap)-1 && height >= heightmap[i+1][j] {
				continue
			}
			lp = append(lp, point{row: i, col: j})
		}
	}
	return lp
}

func part1(heightmap [][]int) int {
	risks := 0
	lowPoints := lowPoints(heightmap)
	for _, p := range lowPoints {
		risks += heightmap[p.row][p.col] + 1
	}
	return risks
}

func exploreBasin(heightmap [][]int, row, col int) (size int) {
	if h := heightmap[row][col]; h == -1 || h == 9 {
		return
	}
	heightmap[row][col] = -1
	size = 1

	// up
	if row-1 >= 0 {
		size += exploreBasin(heightmap, row-1, col)
	}
	// right
	if col+1 < len(heightmap[0]) {
		size += exploreBasin(heightmap, row, col+1)
	}
	// down
	if row+1 < len(heightmap) {
		size += exploreBasin(heightmap, row+1, col)
	}
	// left
	if col-1 >= 0 {
		size += exploreBasin(heightmap, row, col-1)
	}
	return size
}

func part2(heightmap [][]int) int {
	lowPoints := lowPoints(heightmap)
	sizes := make([]int, 0, len(lowPoints))
	for _, p := range lowPoints {
		sizes = append(sizes, exploreBasin(heightmap, p.row, p.col))
	}
	sort.Ints(sizes)
	largest3 := sizes[len(sizes)-3:]
	return largest3[0] * largest3[1] * largest3[2]
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\tinput : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(sample))
	fmt.Printf("\tinput : %d\n", part2(input))

	// Part One
	// 	sample: 15
	// 	input : 545
	// Part Two
	// 	sample: 1134
	// 	input : 950600
}
