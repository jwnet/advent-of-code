package main

import (
	"fmt"
	"log"
	"strconv"
)

func binaryToDecimal(b string) int {
	i, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		log.Fatalln("failed to convert binary to decimal")
	}
	return int(i)
}

// part one
func calculatePower(data []string) int {
	var zeros = make([]int, len(data[0]))
	for _, bnum := range data {
		for i, bit := range bnum {
			if bit == '0' {
				zeros[i]++
			}
		}
	}

	epsilon, gamma := "", ""
	for _, count := range zeros {
		if count > (len(data) / 2) {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	e := binaryToDecimal(epsilon)
	g := binaryToDecimal(gamma)
	return e * g
}

func countBitOccrsByCol(binNums []string) [][2]int {
	nbits := len(binNums[0])
	occrs := make([][2]int, nbits)
	for _, num := range binNums {
		for i, bit := range num {
			switch bit {
			case '0':
				occrs[i][0]++
			case '1':
				occrs[i][1]++
			}
		}
	}
	return occrs
}

func mostCommonBitByCol(binNums []string, tiebreaker rune) []rune {
	mostCommon := make([]rune, len(binNums[0]))
	occrs := countBitOccrsByCol(binNums)
	for i, counts := range occrs {
		switch {
		case counts[0] > counts[1]:
			mostCommon[i] = '0'
		case counts[0] < counts[1]:
			mostCommon[i] = '1'
		default:
			mostCommon[i] = tiebreaker
		}
	}
	return mostCommon
}

func leastCommonBitByCol(binNums []string, tiebreaker rune) []rune {
	leastCommon := make([]rune, len(binNums[0]))
	occrs := countBitOccrsByCol(binNums)
	for i, counts := range occrs {
		switch {
		case counts[0] < counts[1]:
			leastCommon[i] = '0'
		case counts[0] > counts[1]:
			leastCommon[i] = '1'
		default:
			leastCommon[i] = tiebreaker
		}
	}
	return leastCommon
}

func filterColByBit(binNums []string, col int, bit rune) []string {
	filtered := make([]string, 0, len(binNums)/2)
	for _, num := range binNums {
		// works because 0 & 1 runes are only a byte long
		if num[col] == byte(bit) {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

func oxygenGen(binNums []string) int {
	filtered := binNums
	for i := 0; len(filtered) > 1; i++ {
		mostCommon := mostCommonBitByCol(filtered, '1')
		filtered = filterColByBit(filtered, i, mostCommon[i])
	}
	return binaryToDecimal(filtered[0])
}

func co2Scrub(binNums []string) int {
	filtered := binNums
	for i := 0; len(filtered) > 1; i++ {
		leastCommon := leastCommonBitByCol(filtered, '0')
		filtered = filterColByBit(filtered, i, leastCommon[i])
	}
	return binaryToDecimal(filtered[0])
}

// part two
func calculateLifeSupport(binNums []string) int {
	return oxygenGen(binNums) * co2Scrub(binNums)
}

func main() {
	// part one
	fmt.Println(calculatePower(input)) // output: 2003336

	// part two
	fmt.Println(calculateLifeSupport(input)) // output: 1877139
}
