package main

import (
	"fmt"
	"math"
)

type rotation struct {
	x, y, z int
}

var rotations = []rotation{
	{1, 2, 3}, {1, 3, -2}, {1, -2, -3}, {1, -3, 2},
	{-1, 2, -3}, {-1, 3, 2}, {-1, -2, 3}, {-1, -3, -2},
	{2, 1, -3}, {2, 3, 1}, {2, -1, 3}, {2, -3, -1},
	{-2, 1, 3}, {-2, 3, -1}, {-2, -1, -3}, {-2, -3, 1},
	{3, 1, 2}, {3, 2, -1}, {3, -1, -2}, {3, -2, 1},
	{-3, 1, -2}, {-3, 2, 1}, {-3, -1, 2}, {-3, -2, -1},
}

type coord struct {
	x, y, z int
}

func (c coord) String() string {
	return fmt.Sprintf("%5d, %5d, %5d", c.x, c.y, c.z)
}

type probe struct {
	c, oc coord
}

func (p probe) String() string {
	return fmt.Sprintf("{x: %4d, y: %4d, z: %4d}", p.c.x, p.c.y, p.c.z)
}

func (p *probe) rotate(rot int) {
	if rot < 0 || rot >= len(rotations) {
		println("coord.rotate: out of bounds")
		return
	}
	p.c.x = p.transform(rotations[rot].x)
	p.c.y = p.transform(rotations[rot].y)
	p.c.z = p.transform(rotations[rot].z)
}

func (p probe) transform(r int) int {
	n := 0
	switch r {
	case 1:
		n = p.oc.x
	case -1:
		n = -p.oc.x
	case 2:
		n = p.oc.y
	case -2:
		n = -p.oc.y
	case 3:
		n = p.oc.z
	case -3:
		n = -p.oc.z
	}
	return n
}

type scanner struct {
	probes []*probe
	rot    int   // index of current rotation
	offset coord // offset from main list of probes
}

func (s *scanner) cycleRotate() {
	s.rot = (s.rot + 1) % len(rotations)
	for _, pr := range s.probes {
		pr.rotate(s.rot)
	}
}

var scannerCoords = []coord{{0, 0, 0}}

func match12(s1, s2 scanner) (offset coord, ok bool) {
	for _, pr1 := range s1.probes {
		pr1_offset := pr1.c
		for _, pr2 := range s2.probes {
			pr2_offset := pr2.c
			if sameProbe(s1, s2, pr1_offset, pr2_offset) {
				ok = true
				offset.x = s1.offset.x + (pr1.c.x - pr2.c.x)
				offset.y = s1.offset.y + (pr1.c.y - pr2.c.y)
				offset.z = s1.offset.z + (pr1.c.z - pr2.c.z)
				scannerCoords = append(scannerCoords, offset)
				return
			}
		}
	}
	return
}

func sameProbe(s1, s2 scanner, offset1, offset2 coord) bool {
	matches := 0
	for _, pr1 := range s1.probes {
		for _, pr2 := range s2.probes {
			if eq(pr1.c, pr2.c, offset1, offset2) {
				matches += 1
			}
		}
		if matches >= 12 {
			return true
		}
	}
	return false
}

func eq(c1, c2, offset1, offset2 coord) bool {
	return c1.x-offset1.x == c2.x-offset2.x &&
		c1.y-offset1.y == c2.y-offset2.y &&
		c1.z-offset1.z == c2.z-offset2.z
}

func addProbes(to map[coord]bool, from scanner, offset coord) {
	for _, pr := range from.probes {
		c := coord{x: pr.c.x + offset.x, y: pr.c.y + offset.y, z: pr.c.z + offset.z}
		to[c] = true
	}
}

func part1(scanners []scanner) int {
	scannerCoords = []coord{{0, 0, 0}}
	allProbes := map[coord]bool{}

	// scanner 0 is our baseline
	for _, p := range scanners[0].probes {
		allProbes[p.oc] = true
	}
	matchedScanners := []scanner{scanners[0]}
	scanners = scanners[1:]
OUTER:
	for {
		for _, cur := range matchedScanners {
			for idx, s2 := range scanners {
				for i := 0; i < len(rotations); i++ {
					if offset, ok := match12(cur, s2); ok {
						addProbes(allProbes, s2, offset)
						s2.offset = offset
						ns := []scanner{s2}
						matchedScanners = append(ns, matchedScanners...)
						scanners = append(scanners[:idx], scanners[idx+1:]...)
						continue OUTER
					}
					s2.cycleRotate()
				}
			}
		}
		break
	}
	return len(allProbes)
}

func part2() int {
	max := 0
	for i, d1 := range scannerCoords {
		for j, d2 := range scannerCoords {
			if i == j {
				continue
			}
			dx := int(math.Abs(float64(d1.x - d2.x)))
			dy := int(math.Abs(float64(d1.y - d2.y)))
			dz := int(math.Abs(float64(d1.z - d2.z)))
			md := dx + dy + dz
			if md > max {
				max = md
			}
		}
	}
	return max
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	p2sample := part2()
	fmt.Printf("\t input: %d\n", part1(input))
	p2input := part2()
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", p2sample)
	fmt.Printf("\t input: %d\n", p2input)

	// Part One
	//     sample: 79
	//      input: 465
	// Part Two
	//     sample: 3621
	//      input: 12149
}
