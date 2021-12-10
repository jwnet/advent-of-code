package main

import (
	"fmt"
	"sort"
)

/*
	Is there not a nice stack in the std library?
	I can't wait for generics.
*/

type runeStack struct {
	root *element
}

type element struct {
	val  rune
	next *element
}

func (s *runeStack) push(r rune) {
	s.root = &element{val: r, next: s.root}
}

func (s *runeStack) pop() (r rune, ok bool) {
	if s.root == nil {
		return
	}
	ok = true
	r = s.root.val
	s.root = s.root.next
	return
}

func part1(data []string) int {
	var points = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	var match = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}

	score := 0
NEXTLINE:
	for _, line := range data {
		stack := runeStack{}
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack.push(r)
			default:
				if m, ok := stack.pop(); !ok || match[r] != m {
					score += points[r]
					continue NEXTLINE
				}
			}
		}
	}
	return score
}

func part2(data []string) int {
	var points = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	var match = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	var complete = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	scores := []int{}
NEXTLINE:
	for _, line := range data {
		stack := runeStack{}
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack.push(r)
			default:
				if m, ok := stack.pop(); !ok || match[r] != m {
					continue NEXTLINE
				}
			}
		}
		score := 0
		for m, ok := stack.pop(); ok; m, ok = stack.pop() {
			score *= 5
			score += points[complete[m]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("	sample: %d\n", part1(sample))
	fmt.Printf("	input : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("	sample: %d\n", part2(sample))
	fmt.Printf("	input : %d\n", part2(input))

	// Part One
	// 	sample: 26397
	// 	input : 341823
	// Part Two
	// 	sample: 288957
	// 	input : 2801302861
}
