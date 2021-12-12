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

func (s signal) decoderFunc() func(string) string {
	var bit4, bit7 uint8
	for _, p := range s.patterns {
		switch len(p) {
		case 3:
			bit7 = bitEncode(p)
		case 4:
			bit4 = bitEncode(p)
		}
	}
	sureThings := map[int]string{5: "1", 6: "7", 7: "4", 10: "3", 11: "9", 12: "5", 13: "0", 15: "6"}
	return func(encoded string) string {
		bitd := bitEncode(encoded)
		digit, ok := sureThings[bits.OnesCount8(bitd^bit4)+bits.OnesCount8(bitd^bit7)+len(encoded)]
		if !ok {
			switch len(encoded) {
			case 5:
				digit = "2"
			case 7:
				digit = "8"
			}
		}
		return digit
	}
}

func (s signal) decodeOutput() int {
	decoder := s.decoderFunc()
	numstr := ""
	for _, d := range s.output {
		numstr += decoder(d)
	}
	num, _ := strconv.Atoi(numstr)
	return num
}

func part2(data []signal) int {
	sum := 0
	for _, s := range data {
		sum += s.decodeOutput()
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
