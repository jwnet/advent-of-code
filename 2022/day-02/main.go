package main

import "fmt"

const (
	ROCK     int = 1
	PAPER    int = 2
	SCISSORS int = 3
)

// part 1: 14827
// part 2: 13889
func main() {
	totalScore := 0
	for _, round := range input {
		op, you := toRPS(round[0]), toRPS(round[1])
		totalScore += matchScore(op, you)
	}
	fmt.Printf("part 1: %d\n", totalScore)

	totalScore2 := 0
	for _, round := range input {
		op := toRPS(round[0])
		totalScore2 += matchScore(op, chooseMove(op, round[1]))
	}
	fmt.Printf("part 2: %d\n", totalScore2)

}

func toRPS(r rune) int {
	switch r {
	case 'A', 'X':
		return ROCK
	case 'B', 'Y':
		return PAPER
	case 'C', 'Z':
		return SCISSORS
	}
	panic("bad input")
}

func chooseMove(op int, you rune) int {
	switch you {
	case 'X': // lose
		switch op {
		case ROCK:
			return SCISSORS
		case PAPER:
			return ROCK
		case SCISSORS:
			return PAPER
		}
	case 'Y': // draw
		switch op {
		case ROCK:
			return ROCK
		case PAPER:
			return PAPER
		case SCISSORS:
			return SCISSORS
		}
	case 'Z': // win
		switch op {
		case ROCK:
			return PAPER
		case PAPER:
			return SCISSORS
		case SCISSORS:
			return ROCK
		}
	}
	panic("bad input")
}

func matchScore(op, you int) int {
	score := you
	switch op {
	case ROCK:
		switch you {
		case ROCK:
			score += 3
		case PAPER:
			score += 6
		case SCISSORS:
			score += 0
		}
	case PAPER:
		switch you {
		case ROCK:
			score += 0
		case PAPER:
			score += 3
		case SCISSORS:
			score += 6
		}
	case SCISSORS:
		switch you {
		case ROCK:
			score += 6
		case PAPER:
			score += 0
		case SCISSORS:
			score += 3
		}
	}
	return score
}
