package main

import (
	"fmt"
	"strings"
)

func highestScore(rows []string) int {
	nrows := len(rows)
	ncols := len(rows[0])
	high := 0

	up := func(h byte, i, j int) int {
		ntrees := 0
		for x := i - 1; x >= 0; x-- {
			ntrees++
			if rows[x][j] >= h {
				break
			}
		}
		return ntrees
	}
	down := func(h byte, i, j int) int {
		ntrees := 0
		for x := i + 1; x < nrows; x++ {
			ntrees++
			if rows[x][j] >= h {
				break
			}
		}
		return ntrees
	}
	left := func(h byte, i, j int) int {
		ntrees := 0
		for x := j - 1; x >= 0; x-- {
			ntrees++
			if rows[i][x] >= h {
				break
			}
		}
		return ntrees
	}
	right := func(h byte, i, j int) int {
		ntrees := 0
		for x := j + 1; x < ncols; x++ {
			ntrees++
			if rows[i][x] >= h {
				break
			}
		}
		return ntrees
	}

	for i := 1; i < nrows-1; i++ {
		for j := 1; j < ncols-1; j++ {
			height := rows[i][j]
			scores := [4]int{up(height, i, j), down(height, i, j), left(height, i, j), right(height, i, j)}
			cur := 1
			for _, score := range scores {
				if score > 0 {
					cur *= score
				}
			}
			if cur > high {
				high = cur
			}
		}
	}
	return high
}

func countVisible(rows []string) int {
	nrows := len(rows)
	ncols := len(rows[0])

	// outer trees
	visible := 2*(nrows+ncols) - 4

	up := func(h byte, i, j int) bool {
		for x := i - 1; x >= 0; x-- {
			if rows[x][j] >= h {
				return false
			}
		}
		return true
	}
	down := func(h byte, i, j int) bool {
		for x := i + 1; x < nrows; x++ {
			if rows[x][j] >= h {
				return false
			}
		}
		return true
	}
	left := func(h byte, i, j int) bool {
		for x := j - 1; x >= 0; x-- {
			if rows[i][x] >= h {
				return false
			}
		}
		return true
	}
	right := func(h byte, i, j int) bool {
		for x := j + 1; x < ncols; x++ {
			if rows[i][x] >= h {
				return false
			}
		}
		return true
	}

	for i := 1; i < nrows-1; i++ {
		for j := 1; j < ncols-1; j++ {
			height := rows[i][j]
			if up(height, i, j) || down(height, i, j) || left(height, i, j) || right(height, i, j) {
				visible++
			}
		}
	}
	return visible
}

// part 1: 1711
// part 2: 301392
func main() {
	rows := strings.Split(input, "\n")

	fmt.Printf("part 1: %d\n", countVisible(rows))
	fmt.Printf("part 2: %d\n", highestScore(rows))
}
