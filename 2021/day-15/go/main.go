package main

import (
	"fmt"
)

func expandCave(risks [][]int) [][]int {
	m, n := size2d(risks)
	expanded := make2d[int](m*5, n*5)
	for i, row := range expanded {
		for j := range row {
			risk := 0
			switch {
			case i < m && j < n:
				risk = risks[i][j]
			case i < m:
				risk = expanded[i][j-n] + 1 
			case j < n :
				risk = expanded[i-m][j] + 1
			default:
				risk = expanded[i-m][j-n] + 2
			}
			if risk > 9 {
				risk = risk % 9
			}
			expanded[i][j] = risk
		}
	}
	return expanded
}

func updateLowestRunningRisk(rpq *rpqueue, rr int, p point) {
	if rp, ok := rpq.get(p); ok {
		if rp.r + rr < rp.rr {
			rpq.update(rp.p, rp.r + rr)
		}
	}
}

func lowestRiskPath(risks [][]int) int {
	m, n := size2d(risks)
	if m == 0 || n == 0 {
		return 0
	}

	rps := make([]*riskPoint, m*n)
	for r, row := range risks {
		for c, risk := range row {
			rps[r*m+c] = &riskPoint{p: point{r,c}, r: risk, rr: 10000000}
		}
	}
	rps[0].rr = 0
	endPoint := rps[len(rps)-1]

	rpq := rpqueueFromOrdered(rps)
	for rp, ok := rpq.pop(); ok; rp, ok = rpq.pop() {
		updateLowestRunningRisk(rpq, rp.rr, point{rp.p.r-1, rp.p.c})
		updateLowestRunningRisk(rpq, rp.rr, point{rp.p.r+1, rp.p.c})
		updateLowestRunningRisk(rpq, rp.rr, point{rp.p.r, rp.p.c-1})
		updateLowestRunningRisk(rpq, rp.rr, point{rp.p.r, rp.p.c+1})
	}
	return endPoint.rr
}

func part1(risks [][]int) int {
	return lowestRiskPath(risks)
}

func part2(risks [][]int) int {
	expandedRisks := expandCave(risks)
	return lowestRiskPath(expandedRisks)
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("	sample: %d\n", part1(sample))
	fmt.Printf("	input : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("	sample: %d\n", part2(sample))
	fmt.Printf("	input : %d\n", part2(input))

	// Part One
	// 	sample: 40
	// 	input : 423
	// Part Two
	// 	sample: 315
	// 	input : 2778
}
