package main

import (
	"fmt"
)

func isLowercase(s string) bool {
	for _, r := range s {
		if !(r >= 'a' && r <= 'z') {
			return false
		}
	}
	return true
}

func noDoubleVisits(visits map[string]uint8) bool {
	for _, nvisits := range visits {
		if nvisits >= 2 {
			return false
		}
	}
	return true
}

type caveSystem struct {
	connections map[string]*cave
}

func caveSystemFrom(data []connectionData) caveSystem {
	g := caveSystem{connections: make(map[string]*cave)}
	for _, cd := range data {
		g.addConnection(cd)
	}
	return g
}

func (p *caveSystem) addConnection(cd connectionData) {
	na, found := p.connections[cd.a]
	if !found {
		na = &cave{label: cd.a, small: isLowercase(cd.a), connections: make(map[string]*cave)}
		p.connections[cd.a] = na
	}
	nb, found := p.connections[cd.b]
	if !found {
		nb = &cave{label: cd.b, small: isLowercase(cd.b), connections: make(map[string]*cave)}
		p.connections[cd.b] = nb
	}
	na.addConnection(nb)
	nb.addConnection(na)
}

func (cs *caveSystem) nPathsToEnd(oneSmallCaveDoubleVisit bool) int {
	start, foundStart := cs.connections["start"]
	_, foundEnd := cs.connections["end"]
	if !(foundStart && foundEnd) {
		return 0
	}

	smallCaveVisits := make(map[string]uint8)
	return start.nPathsToEnd(oneSmallCaveDoubleVisit, smallCaveVisits)
}

type cave struct {
	label       string
	small       bool
	connections map[string]*cave
}

func (nd *cave) addConnection(n *cave) {
	nd.connections[n.label] = n
}

func (cv *cave) nPathsToEnd(oneSmallCaveDoubleVisit bool, smallCaveVisits map[string]uint8) int {
	if cv.small {
		smallCaveVisits[cv.label] += 1
		defer func() {
			smallCaveVisits[cv.label]--
		}()
	}

	allowedSmallCaveVisits := uint8(1)
	if oneSmallCaveDoubleVisit && noDoubleVisits(smallCaveVisits) {
		allowedSmallCaveVisits = 2
	}

	npaths := 0
	for label, n := range cv.connections {
		switch {
		case label == "start":
			continue
		case label == "end":
			npaths += 1
			continue
		default:
			if nvisits, ok := smallCaveVisits[label]; ok && nvisits >= allowedSmallCaveVisits {
				continue
			}
		}
		npaths += n.nPathsToEnd(oneSmallCaveDoubleVisit, smallCaveVisits)
	}
	return npaths
}

func part1(data []connectionData) int {
	g := caveSystemFrom(data)
	return g.nPathsToEnd(false)
}

func part2(data []connectionData) int {
	g := caveSystemFrom(data)
	return g.nPathsToEnd(true)
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample1: %d\n", part1(sample1))
	fmt.Printf("\tsample2: %d\n", part1(sample2))
	fmt.Printf("\tsample3: %d\n", part1(sample3))
	fmt.Printf("\t  input: %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample1: %d\n", part2(sample1))
	fmt.Printf("\tsample2: %d\n", part2(sample2))
	fmt.Printf("\tsample3: %d\n", part2(sample3))
	fmt.Printf("\t  input: %d\n", part2(input))

	// 	Part One
	// 		sample1: 10
	// 		sample2: 19
	// 		sample3: 226
	// 		  input: 5576
	// 	Part Two
	// 		sample1: 36
	// 		sample2: 103
	// 		sample3: 3509
	// 		  input: 152837
}
