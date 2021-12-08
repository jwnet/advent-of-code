package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type runeSorter []rune

func (r runeSorter) Less(i, j int) bool {
	return r[i] < r[j]
}
func (r runeSorter) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r runeSorter) Len() int {
	return len(r)
}
func sortrunes(s string) string {
	r := runeSorter([]rune(s))
	sort.Sort(r)
	return string(r)
}

func diff(s1, s2 string) int {
	uniq := 0
	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			uniq++
		}
	}
	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			uniq++
		}
	}
	return uniq
}

func decodePatterns(patterns []string) map[string]string {
	decoder := make(map[string]string, 10)
	helper := make(map[int]string, 4)
	remaining := make([]string, 0, 6)
	for _, p := range patterns {
		p = sortrunes(p)
		switch len(p) {
		case 2:
			decoder[p] = "1"
			helper[1] = p
		case 3:
			decoder[p] = "7"
			helper[7] = p
		case 4:
			decoder[p] = "4"
			helper[4] = p
		case 7:
			decoder[p] = "8"
			helper[8] = p
		default:
			remaining = append(remaining, p)
		}
	}
	for _, p := range remaining {
		p = sortrunes(p)
		switch len(p) {
		case 5: // 2, 3, 5
			switch diff(p, helper[4]) {
			case 5:
				decoder[p] = "2"
			case 3:
				if diff(p, helper[7]) == 4 {
					decoder[p] = "5"
				} else {
					decoder[p] = "3"
				}
			}
		case 6: // 0, 6, 9
			switch diff(p, helper[4]) {
			case 2:
				decoder[p] = "9"
			case 4:
				if diff(p, helper[7]) == 3 {
					decoder[p] = "0"
				} else {
					decoder[p] = "6"
				}
			}
		}
	}
	return decoder
}

func part1(data []signal) int {
	count1478 := 0
	for _, s := range data {
		for _, n := range s.output {
			switch len(n) {
			case 2, 3, 4, 7:
				count1478++
			}
		}
	}
	return count1478
}

func part2(data []signal) int {
	sum := 0
	for _, s := range data {
		decoder := decodePatterns(s.patterns)
		numstr := ""
		for _, n := range s.output {
			n = sortrunes(n)
			numstr += decoder[n]
		}
		num, err := strconv.Atoi(numstr)
		if err != nil {
			log.Fatalln(err)
		}
		sum += num
	}
	return sum
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\tinput : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(sample))
	fmt.Printf("\tinput : %d\n", part2(input))

	// Part One
	// 	sample: 26
	// 	input : 495
	// Part Two
	// 	sample: 61229
	// 	input : 1055164
}
