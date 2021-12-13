package main

import (
	"fmt"
	"strings"
)

func foldDot(dot point, fold point) (d point, folded bool) {
	if fold.x > 0 && dot.x > fold.x {
		return point{y: dot.y, x: fold.x - (dot.x - fold.x)}, true
	}
	if fold.y > 0 && dot.y > fold.y {
		return point{y: fold.y - (dot.y - fold.y), x: dot.x}, true
	}
	return dot, false
}

func printDotMap(dots map[point]uint8) string {
	xmax, ymax := 0, 0
	for d := range dots {
		if d.x > xmax {
			xmax = d.x
		}
		if d.y > ymax {
			ymax = d.y
		}
	}
	grid := make([][]rune, ymax+1)
	for y := range grid {
		grid[y] = make([]rune, xmax+1)
	}

	for d := range dots {
		grid[d.y][d.x] = '#'
	}

	b := strings.Builder{}
	for _, row := range grid {
		for _, val := range row {
			if val == 0 {
				b.WriteRune(' ')
			} else {
				b.WriteRune(val)
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func part1(data manual) int {
	firstFold := data.folds[0]
	dotMap := make(map[point]uint8)
	for _, d := range data.dots {
		fd, _ := foldDot(d, firstFold)
		dotMap[fd] += 1
	}
	return len(dotMap)
}

func part2(data manual) string {
	dotMap := make(map[point]uint8, len(data.dots))
	for _, d := range data.dots {
		dotMap[d] += 1
	}

	for _, f := range data.folds {
		for d := range dotMap {
			if fd, folded := foldDot(d, f); folded {
				delete(dotMap, d)
				dotMap[fd] += 1
			}
		}
	}
	return printDotMap(dotMap)
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\tinput : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample:\n\n%s\n", part2(sample))
	fmt.Printf("\tinput :\n\n%s\n", part2(input)) // LRFJBJEH

	// Part One
	// 	sample: 17
	// 	input : 706
	// Part Two
	// 	sample:

	// #####
	// #   #
	// #   #
	// #   #
	// #####

	// 	input :

	// #    ###  ####   ## ###    ## #### #  #
	// #    #  # #       # #  #    # #    #  #
	// #    #  # ###     # ###     # ###  ####
	// #    ###  #       # #  #    # #    #  #
	// #    # #  #    #  # #  # #  # #    #  #
	// #### #  # #     ##  ###   ##  #### #  #

}
