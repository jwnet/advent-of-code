package main

import "fmt"

type point struct {
	r, c int
}

/*
	Is there no good way to copy 2d arrays???
*/

func copy2d(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for r, row := range src {
		dst[r] = make([]int, len(row))
		for c, v := range row {
			dst[r][c] = v
		}
	}
	return dst
}

func propagateEnergy(data [][]int, r, c int) (flashes int) {
	// check out of bounds
	if r < 0 || r >= len(data) || c < 0 || c >= len(data[0]) {
		return
	}
	// already flashed
	if data[r][c] == 0 {
		return
	}
	// increase energy
	if data[r][c] < 9 {
		data[r][c]++
		return
	}

	// FLASH!!!
	flashes += 1
	data[r][c] = 0

	flashes += propagateEnergy(data, r-1, c-1)
	flashes += propagateEnergy(data, r-1, c)
	flashes += propagateEnergy(data, r-1, c+1)
	flashes += propagateEnergy(data, r, c-1)
	flashes += propagateEnergy(data, r, c+1)
	flashes += propagateEnergy(data, r+1, c-1)
	flashes += propagateEnergy(data, r+1, c)
	flashes += propagateEnergy(data, r+1, c+1)

	return flashes
}

func step(data [][]int) (flashes int) {
	flashes = 0
	flashPoints := []point{}
	for r, energyLevels := range data {
		for c, e := range energyLevels {
			if e == 9 {
				flashPoints = append(flashPoints, point{r, c})
			}
			data[r][c]++
		}
	}
	for _, p := range flashPoints {
		flashes += propagateEnergy(data, p.r, p.c)
	}
	return flashes
}

func part1(data [][]int) int {
	grid := copy2d(data)
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += step(grid)
	}
	return flashes
}

func part2(data [][]int) int {
	grid := copy2d(data)
	steps := 1
	for {
		if flashes := step(grid); flashes == len(grid)*len(grid[0]) {
			return steps
		}
		steps++
	}
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("	sample: %d\n", part1(sample))
	fmt.Printf("	input : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("	sample: %d\n", part2(sample))
	fmt.Printf("	input : %d\n", part2(input))

	// Part One
	// 	sample: 1656
	// 	input : 1755
	// Part Two
	// 	sample: 195
	// 	input : 212

}
