package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jwnet/useful/list"
)

type instruction struct {
	from, to, n int
}

func stacksFrom(s string) []*list.List[rune] {
	s = strings.ReplaceAll(s, "[", " ")
	s = strings.ReplaceAll(s, "]", " ")
	s = strings.ReplaceAll(s, "    ", "   .")
	s = strings.ReplaceAll(s, "   ", " ")
	s = strings.ReplaceAll(s, " ", "")

	lines := strings.Split(s, "\n")
	var stacks []*list.List[rune]

	labels := lines[len(lines)-1]
	for range labels {
		stacks = append(stacks, &list.List[rune]{})
	}

	for i := len(lines) - 2; i >= 0; i-- {
		for j, r := range lines[i] {
			if r == '.' {
				continue
			}
			stacks[j].Push(r)
		}
	}

	return stacks
}

func instructionsFrom(s string) []instruction {
	var instructions []instruction

	s = strings.ReplaceAll(s, "move ", "")
	s = strings.ReplaceAll(s, "from ", "")
	s = strings.ReplaceAll(s, "to ", "")

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		nums := strings.Split(line, " ")
		n, err := strconv.Atoi(nums[0])
		from, err := strconv.Atoi(nums[1])
		to, err := strconv.Atoi(nums[2])

		if err != nil {
			panic("impossible!")
		}

		instructions = append(instructions, instruction{from: from - 1, to: to - 1, n: n})
	}
	return instructions
}

func stackTops(stacks []*list.List[rune]) string {
	s := strings.Builder{}
	for _, stack := range stacks {
		_, r := stack.Front()
		s.WriteRune(r)
	}
	return s.String()
}

func top2top(s1, s2 *list.List[rune]) {
	ok, r := s1.Pop()
	if !ok {
		println(r)
		panic("oh no")
	}
	s2.Push(r)
}

// part 1: VJSFHWGFT
// part 2: LCTQFBVZV
func main() {
	sections := strings.Split(input, "\n\n")
	startState, instructionSet := sections[0], sections[1]

	stacks := stacksFrom(startState)
	stacks2 := stacksFrom(startState)

	instructions := instructionsFrom(instructionSet)

	for _, ins := range instructions {
		// part 1
		for i := 0; i < ins.n; i++ {
			top2top(stacks[ins.from], stacks[ins.to])
		}
		// part 2
		tempStack := &list.List[rune]{}
		for i := 0; i < ins.n; i++ {
			top2top(stacks2[ins.from], tempStack)
		}
		for i := 0; i < ins.n; i++ {
			top2top(tempStack, stacks2[ins.to])
		}
	}

	fmt.Printf("part 1: %s\n", stackTops(stacks))
	fmt.Printf("part 2: %s\n", stackTops(stacks2))
}
