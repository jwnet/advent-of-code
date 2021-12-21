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

type beacon struct {
	c, oc coord // oc is the original coordinates, and does not get changed in rotation
}

func (b *beacon) rotate(rot int) {
	if rot < 0 || rot >= len(rotations) {
		println("coord.rotate: out of bounds")
		return
	}
	b.c.x = b.transform(rotations[rot].x)
	b.c.y = b.transform(rotations[rot].y)
	b.c.z = b.transform(rotations[rot].z)
}

func (b *beacon) transform(r int) int {
	n := 0
	switch r {
	case 1:
		n = b.oc.x
	case -1:
		n = -b.oc.x
	case 2:
		n = b.oc.y
	case -2:
		n = -b.oc.y
	case 3:
		n = b.oc.z
	case -3:
		n = -b.oc.z
	}
	return n
}

type scanner struct {
	beacons []*beacon
	rot     int   // index of current rotation
	offset  coord // offset from main list of probes--from scanner 0
}

func (s *scanner) cycleRotate() {
	s.rot = (s.rot + 1) % len(rotations)
	for _, b := range s.beacons {
		b.rotate(s.rot)
	}
}

func match12(s1, s2 scanner) (offset coord, ok bool) {
	for _, b1 := range s1.beacons {
		b1_offset := b1.c
		for _, b2 := range s2.beacons {
			b2_offset := b2.c
			if sameBeacon(s1, s2, b1_offset, b2_offset) {
				ok = true
				offset.x = s1.offset.x + (b1.c.x - b2.c.x)
				offset.y = s1.offset.y + (b1.c.y - b2.c.y)
				offset.z = s1.offset.z + (b1.c.z - b2.c.z)
				return
			}
		}
	}
	return
}

func sameBeacon(s1, s2 scanner, offset1, offset2 coord) bool {
	matches := 0
	for _, b1 := range s1.beacons {
		for _, b2 := range s2.beacons {
			if eq(b1.c, b2.c, offset1, offset2) {
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

func addBeacons(to map[coord]bool, s scanner) {
	for _, b := range s.beacons {
		c := coord{x: b.c.x + s.offset.x, y: b.c.y + s.offset.y, z: b.c.z + s.offset.z}
		to[c] = true
	}
}

func part1(scanners []scanner) int {
	scannerCoords = []coord{{0, 0, 0}}
	allBeacons := map[coord]bool{}
	for _, b := range scanners[0].beacons {
		allBeacons[b.c] = true
	}
	matchedScanners := []scanner{scanners[0]}
	scanners = scanners[1:]
OUTER:
	for {
		for _, cur := range matchedScanners {
			for idx, s := range scanners {
				for i := 0; i < len(rotations); i++ {
					if offset, ok := match12(cur, s); ok {
						s.offset = offset
						addBeacons(allBeacons, s)

						// add self to front of matchedScanners slice
						scannerCoords = append(scannerCoords, offset)
						temp := []scanner{s}
						matchedScanners = append(temp, matchedScanners...)

						// remove self from scanner slice
						scanners = append(scanners[:idx], scanners[idx+1:]...)
						continue OUTER
					}
					s.cycleRotate()
				}
			}
		}
		break
	}
	return len(allBeacons)
}

var scannerCoords = []coord{}

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
