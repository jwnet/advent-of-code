package main

import (
	"fmt"
	"sort"
)

type point struct {
	r, c int
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
			lp = append(lp, point{r: i, c: j})
		}
	}
	return lp
}

func part1(heightmap [][]int) int {
	risks := 0
	lowPoints := lowPoints(heightmap)
	for _, p := range lowPoints {
		risks += heightmap[p.r][p.c] + 1
	}
	return risks
}

func exploreBasin(heightmap [][]int, start point, seen map[point]bool) (size int) {
	seen[start] = true
	height := heightmap[start.r][start.c]
	if height == 9 {
		return
	}
	size = 1

	up := point{start.r - 1, start.c}
	if _, visited := seen[up]; !visited && up.r >= 0 {
		size += exploreBasin(heightmap, up, seen)
	}

	right := point{start.r, start.c + 1}
	if _, visited := seen[right]; !visited && right.c < len(heightmap[0]) {
		size += exploreBasin(heightmap, right, seen)
	}

	down := point{start.r + 1, start.c}
	if _, visited := seen[down]; !visited && down.r < len(heightmap) {
		size += exploreBasin(heightmap, down, seen)
	}

	left := point{start.r, start.c - 1}
	if _, visited := seen[left]; !visited && left.c >= 0 {
		size += exploreBasin(heightmap, left, seen)
	}

	return size
}

func part2(heightmap [][]int) int {
	lowPoints := lowPoints(heightmap)
	sizes := make([]int, 0, len(lowPoints))
	for _, p := range lowPoints {
		seen := make(map[point]bool)
		sizes = append(sizes, exploreBasin(heightmap, p, seen))

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
