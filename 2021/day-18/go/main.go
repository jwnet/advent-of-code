package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	for _, s := range sample {
		s.Update()
	}
	for _, s := range input {
		s.Update()
	}
}

type direction uint8
const (
	none direction = iota
	left
	right
)

type snailnum struct {
	ll, rl int
	lp, rp *snailnum
	parent *snailnum
	level int
	side direction
}

func add(a, b snailnum) snailnum {
	sum := snailnum{lp: a.clone(), rp: b.clone()}
	sum.Update()
	sum.Reduce()
	return sum
}

func (s *snailnum) clone() *snailnum {
	cl := new(snailnum)
	cl.level = s.level
	cl.side = s.side
	cl.ll, cl.rl = s.ll, s.rl
	if s.lp != nil {
		cl.lp = s.lp.clone()
		cl.lp.parent = cl
	}
	if s.rp != nil {
		cl.rp = s.rp.clone()
		cl.rp.parent = cl
	}
	return cl
}

func (s *snailnum) Update() {
	s.update(s.level, s.parent)
}

func (s *snailnum) update(level int, parent *snailnum) {
	s.level = level
	s.parent = parent
	if s.lp != nil {
		s.lp.side = left
		s.lp.update(level+1, s)
	}
	if s.rp != nil {
		s.rp.side = right
		s.rp.update(level+1, s)
	}
}

func (s snailnum) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	if s.lp != nil {
		sb.WriteString(s.lp.String())
	} else {
		sb.WriteString(strconv.Itoa(s.ll))
	}
	sb.WriteRune(',')
	if s.rp != nil {
		sb.WriteString(s.rp.String())
	} else {
		sb.WriteString(strconv.Itoa(s.rl))
	}
	sb.WriteRune(']')
	return sb.String()
}

func (s *snailnum) Reduce() {
	for s.reduce() {}
}

func (s *snailnum) reduce() bool {
	if s.explode() {
		return true
	}
	if s.split() {
		return true
	}
	return false
}

func (s *snailnum) split() bool {
	if s.lp == nil && s.ll > 9 {
		half := float64(s.ll) / 2.0
		s.lp = &snailnum{ll: int(math.Floor(half)), rl: int(math.Ceil(half))}
		s.Update()
		return true
	} else if s.lp != nil && s.lp.split() {
		return true
	}
	if s.rp == nil && s.rl > 9 {
		half := float64(s.rl) / 2.0
		s.rp = &snailnum{ll: int(math.Floor(half)), rl: int(math.Ceil(half))}
		s.Update()
		return true
	} else if s.rp != nil && s.rp.split() {
		return true
	}
	return false
}

func (s *snailnum) explode() bool {
	if s.lp != nil && s.lp.level == 4 {
		p := s.lp
		s.lp = nil
		s.ll = 0
		s.parent.sendLeft(p.ll, s.side)
		s.sendRight(p.rl, left)
		return true
	}
	if s.rp != nil && s.rp.level == 4 {
		p := s.rp
		s.rp = nil
		s.rl = 0
		s.parent.sendRight(p.rl, s.side)
		s.sendLeft(p.ll, right)
		return true
	}
	if s.lp != nil && s.lp.explode() {
		return true
	}
	if s.rp != nil && s.rp.explode() {
		return true
	}

	return false
}

func (s *snailnum) sendLeft(n int, from direction) {
	switch from {
	case left:
		if s.parent != nil {
			s.parent.sendLeft(n, s.side)
		}
	case right:
		if s.lp == nil {
			s.ll += n
		} else {
			s.lp.sendDownRight(n) 
		}
	}
}

func (s *snailnum) sendRight(n int, from direction) {
	switch from {
	case right:
		if s.parent != nil {
			s.parent.sendRight(n, s.side)
		}
	case left:
		if s.rp == nil {
			s.rl += n
		} else {
			s.rp.sendDownLeft(n) 
		}
	}
}

func (s *snailnum) sendDownRight(n int) {
	if s.rp == nil {
		s.rl += n
	} else {
		s.rp.sendDownRight(n)
	}
}

func (s *snailnum) sendDownLeft(n int) {
	if s.lp == nil {
		s.ll += n
	} else {
		s.lp.sendDownLeft(n)
	}
}

func (s *snailnum) magnitude() int {
	magnitude := 0
	if s.lp == nil {
		magnitude += 3 * s.ll
	} else {
		magnitude += 3 * s.lp.magnitude()
	}
	if s.rp == nil {
		magnitude += 2 * s.rl
	} else {
		magnitude += 2 * s.rp.magnitude()
	}
	return magnitude
}

func part1(nums []snailnum) int {
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		num = add(num, nums[i])
	}
	return num.magnitude()
}

func part2(nums []snailnum) int {
	maxmag := 0
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}
			sum := add(a, b)
			if m := sum.magnitude(); m > maxmag {
				maxmag = m
			}
		}
	}
	return maxmag
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\t input: %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(sample))
	fmt.Printf("\t input: %d\n", part2(input))

	// Part One
	// 	sample: 4140
	// 	 input: 4057
	// Part Two
	// 	sample: 3993
	// 	 input: 4683
}
