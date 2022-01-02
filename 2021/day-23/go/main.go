package main

import (
	"fmt"
	"math"
)

const (
	A = iota
	B
	C
	D
	Empty
)

const (
	RA = 2
	RB = 4
	RC = 6
	RD = 8
)

var rooms = map[int]int{A: RA, B: RB, C: RC, D: RD}
var atype = map[rune]int{'A':A, 'B':B, 'C':C, 'D':D, 0:Empty}
var cost = map[int]int{A: 1, B: 10, C: 100, D: 1000}
var costs []int

type burrow struct {
	hall  [11]rune
	rooms [4][2]rune
	cost  int
}

func (b burrow) String() string {
	s := ""
	for _, r := range b.hall {
		if r == 0 {
			s += " "
		} else {
			s += string(r)
		}
	}
	s += "\n"
	r1, r2 := "##", "##"
	for _, r := range b.rooms {
		if r[0] != 0 {
			r1 = r1 + string(r[0])
		} else {
			r1 = r1 + " "
		}
		if r[1] != 0 {
			r2 = r2 + string(r[1])
		} else {
			r2 = r2 + " "
		}
		r1 = r1 + "#"
		r2 = r2 + "#"
	}
	s += r1 + "#\n" + r2  + "#"
	return s
}

func (b burrow) done() bool {
	return b.rooms[0][0] == 'A' && b.rooms[0][1] == 'A' &&
		b.rooms[1][0] == 'B' && b.rooms[1][1] == 'B' &&
		b.rooms[2][0] == 'C' && b.rooms[2][1] == 'C' &&
		b.rooms[3][0] == 'D' && b.rooms[3][1] == 'D'
}

var hallSpots = [7]int{0,1,3,5,7,9,10}
func arrange(b burrow) {
	if b.done() {
		costs = append(costs, b.cost)
		return
	}
	// make every possible move
	for i, r := range b.hall {
		a := atype[r]
		if a != Empty {
			if nb, ok := b.hallToRoom(i); ok {
				arrange(nb)
			}
		}
	}
	for i, room := range b.rooms {
		A1, A2 := room[0], room[1]
		a1, _ := atype[A1]
		a2, _ := atype[A2]
		if (a1 == Empty && a2 == Empty) || (a1 == i && a2 == i)  {
			continue
		}
		switch {
		case a1 == i && a2 != i:
			for _, hallIdx := range hallSpots {
				if nb, ok := b.roomToHall(i,0,hallIdx); ok {
					arrange(nb)
				}
			}
		case a1 != Empty && a1 != i:
			if nb, ok := b.roomToRoom(i,0); ok {
				arrange(nb)
				continue // because there is no cheaper cost
			}
			for _, hallIdx := range hallSpots {
				if nb, ok := b.roomToHall(i,0,hallIdx); ok {
					arrange(nb)
				}
			}
		case a1 == Empty && a2 != i:
			if nb, ok := b.roomToRoom(i,1); ok {
				arrange(nb)
				continue // because there is no cheaper cost
			}
			for _, hallIdx := range hallSpots {
				if nb, ok := b.roomToHall(i,1,hallIdx); ok {
					arrange(nb)
				}
			}
		}
	}
}

func (b burrow) roomToHall(r, idx, hallIdx int) (nb burrow, ok bool) {
	A := b.rooms[r][idx]
	a := atype[A]
	cri := rooms[r]
	if !b.hallClear(cri, hallIdx) {
		return
	}
	b.rooms[r][idx] = 0
	b.hall[hallIdx] = A
	b.cost += (abs(cri-hallIdx) + 1 + idx) *  cost[a]
	ok = true
	nb = b
	return
}

func (b burrow) roomToRoom(r, idx int) (nb burrow, ok bool) {
	A := b.rooms[r][idx]
	a := atype[A]
	cri := rooms[r]
	ri := rooms[a]
	if !b.hallClear(cri, ri) {
		return
	}
	if b.rooms[a][0] == 0 {
		switch {
		case b.rooms[a][1] == 0:
			b.rooms[r][idx] = 0
			b.rooms[a][1] = A
			b.cost += (abs(cri - ri) + 3 + idx) * cost[a]
			ok = true
		case b.rooms[a][1] == A:
			b.rooms[r][idx] = 0
			b.rooms[a][0] = A
			b.cost += (abs(cri - ri) + 2 + idx) * cost[a]
			ok = true
		}
	}
	nb = b
	return
}

func (b burrow) hallToRoom(idx int) (nb burrow, ok bool) {
	A := b.hall[idx]
	a := atype[A]
	ri, rok := rooms[a]
	if !rok || !b.hallClear(idx, ri) {
		return
	}
	if b.rooms[a][0] == 0 {
		switch {
		case b.rooms[a][1] == 0:
			b.hall[idx] = 0
			b.rooms[a][1] = A
			b.cost += (abs(idx - ri) + 2) * cost[a]
			ok = true
		case b.rooms[a][1] == A:
			b.hall[idx] = 0
			b.rooms[a][0] = A
			b.cost += (abs(idx - ri) + 1) * cost[a]
			ok = true
		}
	}
	nb = b
	return
}

// b.hall[from] is expected to be occupied
func (b burrow) hallClear(from, to int) bool {
	if from < to {
		for i := from+1; i <= to; i++ {
			if b.hall[i] != 0 {
				return false
			}
		}
		return true
	}
	for i := from-1; i >= to; i-- {
		if b.hall[i] != 0 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func part1(b burrow) int {
	costs = []int{}
	arrange(b)
	min := int(1e10)
	for _, cost := range costs {
		if cost < min {
			min = cost
		}
	}
	return min
}

func main() {
	sample := burrow{}
	sample.rooms = [4][2]rune{{'B', 'A'}, {'C', 'D'}, {'B', 'C'}, {'D', 'A'}}
	input := burrow{}
	input.rooms = [4][2]rune{{'B', 'D'}, {'A', 'C'}, {'A', 'B'}, {'D', 'C'}}

	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\t input: %d\n", part1(input))
}
