package main

import (
	"fmt"
	"strings"
)

type sack struct {
	a, b []rune
}

func sackFromString(s string) sack {
	if len(s)%2 != 0 {
		panic("sack's must have an even number of items")
	}

	return sack{a: []rune(s[:len(s)/2]), b: []rune(s[len(s)/2:])}
}

func (s *sack) findCommonItem() rune {
	m := make(map[rune]bool, len(s.a))
	for _, r := range s.a {
		m[r] = true
	}
	for _, r := range s.b {
		if _, ok := m[r]; ok {
			return r
		}
	}
	panic("no duplicate found")
}

func (s *sack) itemMap() map[rune]bool {
	m := make(map[rune]bool, len(s.a)*2)
	for _, r := range s.a {
		m[r] = true
	}
	for _, r := range s.b {
		m[r] = true
	}
	return m
}

func priority(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r-'a') + 1
	}
	if 'A' <= r && r <= 'Z' {
		return int(r-'A') + 27
	}
	panic("invalid item")
}

// I bemoan the lack of sets in the standard library
// part 1: 7793
// part 2: 2499
func main() {
	lines := strings.Split(input, "\n")

	priorities := 0
	for _, line := range lines {
		s := sackFromString(line)
		priorities += priority(s.findCommonItem())
	}
	fmt.Printf("part 1: %d\n", priorities)

	priorities2 := 0
	for i := 0; i < len(lines); i += 3 {
		s1 := sackFromString(lines[i])
		m1 := s1.itemMap()

		common := make(map[rune]bool)
		for _, r := range lines[i+1] {
			if _, ok := m1[r]; ok {
				common[r] = true
			}
		}
		for _, r := range lines[i+2] {
			if _, ok := common[r]; ok {
				priorities2 += priority(r)
				break
			}
		}
	}
	fmt.Printf("part 2: %d\n", priorities2)
}
