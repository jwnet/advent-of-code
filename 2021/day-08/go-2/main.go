package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

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

func bitEncode(p string) uint8 {
	var num uint8 = 0
	for _, r := range p {
		num = num | (1 << (r - 'a'))
	}
	return num
}

func decodeOutput(s signal) int {
	var bit4, bit7 uint8
	decoded := make(map[string]string)
	remaining := make([]string, 0, 6)
	for _, p := range s.patterns {
		switch len(p) {
		case 2:
			decoded[p] = "1"
		case 3:
			decoded[p] = "7"
			bit7 = bitEncode(p)
		case 4:
			decoded[p] = "4"
			bit4 = bitEncode(p)
		case 7:
			decoded[p] = "8"
		default:
			remaining = append(remaining, p)
		}
	}
	for _, p := range remaining {
		bitp := bitEncode(p)
		switch bits.OnesCount8(bitp^bit4) + bits.OnesCount8(bitp^bit7) + len(p) {
		case 10:
			decoded[p] = "3"
		case 11:
			decoded[p] = "9"
		case 12:
			decoded[p] = "5"
		case 13:
			decoded[p] = "0"
		case 14:
			decoded[p] = "2"
		case 15:
			decoded[p] = "6"
		}
	}
	numStr := ""
	for _, d := range s.output {
		numStr += decoded[d]
	}
	num, _ := strconv.Atoi(numStr)
	return num
}

func part2(data []signal) int {
	sum := 0
	for _, s := range data {
		sum += decodeOutput(s)
	}
	return sum
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("	sample: %d\n", part1(sample))
	fmt.Printf("	input : %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("	sample: %d\n", part2(sample))
	fmt.Printf("	input : %d\n", part2(input))

	// Part One
	//     sample: 26
	//     input : 495
	// Part Two
	//     sample: 61229
	//     input : 1055164
}
