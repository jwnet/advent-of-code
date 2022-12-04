package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pair struct {
	a, b assignment
}

func newPair(s string) pair {
	s = strings.Replace(s, ",", "-", 1)
	fields := strings.Split(s, "-")
	var a, b assignment
	var err error
	a.start, err = strconv.Atoi(fields[0])
	a.end, err = strconv.Atoi(fields[1])
	b.start, err = strconv.Atoi(fields[2])
	b.end, err = strconv.Atoi(fields[3])

	if err != nil {
		panic("parse error")
	}
	return pair{a, b}
}

func (p pair) fullOverlap() bool {
	return (p.a.start <= p.b.start && p.b.end <= p.a.end) ||
		(p.b.start <= p.a.start && p.a.end <= p.b.end)
}

func (p pair) overlap() bool {
	return inRange(p.b.start, p.a.start, p.a.end) ||
		inRange(p.b.end, p.a.start, p.a.end) ||
		inRange(p.a.start, p.b.start, p.b.end) ||
		inRange(p.a.end, p.b.start, p.b.end)
}

func inRange(num, start, end int) bool {
	return start <= num && num <= end

}

type assignment struct {
	start, end int
}

// part 1: 511
// part 2: 821
func main() {
	var pairs []pair
	fullOverlap, overlap := 0, 0
	for _, line := range strings.Split(input, "\n") {
		pair := newPair(line)
		pairs = append(pairs, pair)
		if pair.fullOverlap() {
			fullOverlap++
		}
		if pair.overlap() {
			overlap++
		}
	}
	fmt.Printf("part 1: %d\n", fullOverlap)
	fmt.Printf("part 2: %d\n", overlap)
}
