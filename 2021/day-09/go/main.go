package main

import "fmt"

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
			// above
			if i > 0 && height >= heightmap[i-1][j] {
				continue
			}
			// below
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

func basinRight(heightmap [][]int, start point) (size int) {
	size = 0
	for c := start.c + 1; c < len(heightmap[0]); c++ {
		if heightmap[start.r][c] == 9 {
			return size
		}
		size += 1 + basinUp(heightmap, point{start.r, c}) + basinDown(heightmap, point{start.r, c})
	}
	return size
}

func basinLeft(heightmap [][]int, start point) (size int) {
	size = 0
	for c := start.c - 1; c >= 0; c-- {
		if heightmap[start.r][c] == 9 {
			return size
		}
		size += 1 + basinUp(heightmap, point{start.r, c}) + basinDown(heightmap, point{start.r, c})
	}
	return size
}

func basinUp(heightmap [][]int, start point) (size int) {
	size = 0
	for r := start.r - 1; r >= 0; r-- {
		if heightmap[r][start.c] == 9 {
			return size
		}
		size += 1
	}
	return size
}

func basinDown(heightmap [][]int, start point) (size int) {
	size = 0
	for r := start.r + 1; r < len(heightmap); r++ {
		if heightmap[r][start.c] == 9 {
			return size
		}
		size += 1
	}
	return size
}

func part2(heightmap [][]int) int {
	lowPoints := lowPoints(heightmap)
	largest3Basins := []int{-1, -1, -1}
	for _, p := range lowPoints {
		size := 1 + basinLeft(heightmap, p) + basinRight(heightmap, p) + basinDown(heightmap, p) + basinUp(heightmap, p)
		d1 := size - largest3Basins[0]
		d2 := size - largest3Basins[1]
		d3 := size - largest3Basins[2]
		if d1 + d2 + d3 <= 0 {
			continue
		}
		switch {
		case d1 >= d2 && d1 >= d3:
			largest3Basins[0] = size
		case d2 >= d1 && d2 >= d3:
			largest3Basins[1] = size
		case d3 >= d1 && d3 >= d2:
			largest3Basins[2] = size
		}
	}
	return largest3Basins[0] * largest3Basins[1] * largest3Basins[2]
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("  sample: %d\n", part1(sample)) // 15
	fmt.Printf("  input : %d\n", part1(input))  // 545
	fmt.Println("Part Two")
	fmt.Printf("  sample: %d\n", part2(sample))
	// fmt.Printf("  input : %d\n", part2(input))
}
