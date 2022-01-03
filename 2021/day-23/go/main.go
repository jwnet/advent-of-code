package main

import (
	"fmt"
	"math"
)

var hallwayRoomIdxs = [4]int{2, 4, 6, 8}
var roomIdx = map[rune]int{'A': 0, 'B': 1, 'C': 2, 'D': 3}
var cost = map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
var costs []int

type burrow struct {
	hall  [11]rune
	rooms [4][]rune
	cost  int
}

func (b burrow) done() bool {
	for i, room := range b.rooms {
		for _, o := range room {
			if int(o-'A') != i {
				return false
			}
		}
	}
	return true
}

func (b *burrow) clone() burrow {
	nb := burrow{}
	nb.hall = b.hall
	nb.cost = b.cost
	for i, room := range b.rooms {
		nr := make([]rune, len(room))
		copy(nr, room)
		nb.rooms[i] = nr
	}
	return nb
}

// an array of all hallway indexs we can move to
var hallIdxs = [7]int{0, 1, 3, 5, 7, 9, 10}

func arrange(b burrow) {
	if b.done() {
		costs = append(costs, b.cost)
		return
	}
	for i, r := range b.hall {
		if r == 0 {
			continue
		}
		if nb, ok := b.hallToRoom(i); ok {
			arrange(nb)
		}
	}
	for ri, room := range b.rooms {
		for i, occupent := range room {
			if occupent == 0 {
				continue
			}
			if !b.canExitRoom(ri, i) {
				break
			}
			if int(occupent-'A') != ri {
				if nb, ok := b.roomToRoom(ri, i); ok {
					arrange(nb)
					break
				}
			}
			for _, hi := range hallIdxs {
				if nb, ok := b.roomToHall(ri, i, hi); ok {
					arrange(nb)
				}
			}
		}
	}
}

// roomOpen returns whether we can move into a room, and the index
// that we can move into.
func (b *burrow) roomOpen(ri int) (idx int, ok bool) {
	nlen := len(b.rooms[ri])
	for i := nlen - 1; i >= 0; i-- {
		switch b.rooms[ri][i] {
		case rune('A' + ri):
			continue
		case 0:
			return i, true
		default:
			return -1, false
		}
	}
	return -1, false
}

func (b *burrow) canExitRoom(ri, idx int) bool {
	if b.rooms[ri][idx] == 0 {
		return false
	}
	for i := idx - 1; i >= 0; i-- {
		if b.rooms[ri][i] != 0 {
			return false
		}
	}
	if int(b.rooms[ri][idx]-'A') != ri {
		return true
	}
	for i := idx + 1; i < len(b.rooms[ri]); i++ {
		if int(b.rooms[ri][i]-'A') != ri {
			return true
		}
	}
	return false
}

func (b *burrow) roomToHall(r, idx, hallIdx int) (nb burrow, ok bool) {
	occupent := b.rooms[r][idx]
	src_hri := hallwayRoomIdxs[r]
	if !b.hallClear(src_hri, hallIdx) {
		return
	}
	nb = b.clone()
	nb.rooms[r][idx] = 0
	nb.hall[hallIdx] = occupent
	nb.cost += (abs(src_hri-hallIdx) + 1 + idx) * cost[occupent]
	return nb, true
}

func (b *burrow) roomToRoom(r, idx int) (nb burrow, ok bool) {
	occupent := b.rooms[r][idx]
	dst_ri := roomIdx[occupent]
	src_hri := hallwayRoomIdxs[r]
	dst_hri := hallwayRoomIdxs[dst_ri]
	if !b.hallClear(src_hri, dst_hri) {
		return
	}
	oidx, open := b.roomOpen(dst_ri)
	if !open {
		return
	}
	nb = b.clone()
	nb.rooms[r][idx] = 0
	nb.rooms[dst_ri][oidx] = occupent
	nb.cost += (abs(src_hri-dst_hri) + idx + 2 + oidx) * cost[occupent]
	return nb, true
}

func (b *burrow) hallToRoom(idx int) (nb burrow, ok bool) {
	occupent := b.hall[idx]
	dst_ri := roomIdx[occupent]
	src_hri := hallwayRoomIdxs[dst_ri]
	if !b.hallClear(idx, src_hri) {
		return
	}
	oidx, open := b.roomOpen(dst_ri)
	if !open {
		return
	}
	nb = b.clone()
	nb.hall[idx] = 0
	nb.rooms[dst_ri][oidx] = occupent
	nb.cost += (abs(idx-src_hri) + 1 + oidx) * cost[occupent]
	return nb, true
}

// b.hall[from] is expected to be occupied
func (b *burrow) hallClear(from, to int) bool {
	if from < to {
		for i := from + 1; i <= to; i++ {
			if b.hall[i] != 0 {
				return false
			}
		}
		return true
	}
	for i := from - 1; i >= to; i-- {
		if b.hall[i] != 0 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func minCostArrangment(b burrow) int {
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
	sample1 := burrow{}
	sample1.rooms = [4][]rune{{'B', 'A'}, {'C', 'D'}, {'B', 'C'}, {'D', 'A'}}
	input1 := burrow{}
	input1.rooms = [4][]rune{{'B', 'D'}, {'A', 'C'}, {'A', 'B'}, {'D', 'C'}}

	sample2 := burrow{}
	sample2.rooms = [4][]rune{
		{'B', 'D', 'D', 'A'},
		{'C', 'C', 'B', 'D'},
		{'B', 'B', 'A', 'C'},
		{'D', 'A', 'C', 'A'},
	}
	input2 := burrow{}
	input2.rooms = [4][]rune{
		{'B', 'D', 'D', 'D'},
		{'A', 'C', 'B', 'C'},
		{'A', 'B', 'A', 'B'},
		{'D', 'A', 'C', 'C'},
	}

	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", minCostArrangment(sample1))
	fmt.Printf("\t input: %d\n", minCostArrangment(input1))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", minCostArrangment(sample2))
	fmt.Printf("\t input: %d\n", minCostArrangment(input2))

	// Part One
	//	 sample: 12521
	//	  input: 15237
	// Part Two
	//	 sample: 44169
	//	  input: 47509
}
