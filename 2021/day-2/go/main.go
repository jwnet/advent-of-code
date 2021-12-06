package main

import "fmt"

// part one
func calculatePosition(data []movement) int {
	horizontal, depth := 0, 0
	for _, m := range data {
		switch m.direction {
		case "forward":
			horizontal += m.num
		case "up":
			depth -= m.num
		case "down":
			depth += m.num
		}
	}
	return horizontal * depth
}

// part two
func calculateAdvancedPosition(data []movement) int {
	horizontal, aim, depth := 0, 0, 0
	for _, m := range data {
		switch m.direction {
		case "forward":
			horizontal += m.num
			depth += m.num * aim
		case "up":
			aim -= m.num
		case "down":
			aim += m.num
		}
	}
	return horizontal * depth
}

func main() {
	// part one
	fmt.Println(calculatePosition(input)) // output: 1938402

	// part two
	fmt.Println(calculateAdvancedPosition(input)) // output: 1947878632
}
