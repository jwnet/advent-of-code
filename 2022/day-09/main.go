package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jwnet/useful/set"
)

type rope struct {
	knots []coord
}

func (r *rope) head() coord {
	return r.knots[0]
}

func (r *rope) setHead(c coord) {
	r.knots[0] = c
}

func (r *rope) tail() coord {
	return r.knots[len(r.knots)-1]
}

func (r *rope) get(i int) coord {
	return r.knots[i]
}

func (r *rope) set(i int, c coord) {
	r.knots[i] = c
}

type coord struct {
	r, c int
}

func (c coord) String() string {
	return fmt.Sprintf("row: %d col: %d", c.r, c.c)
}

func up(c coord) coord {
	c.r += 1
	return c
}

func down(c coord) coord {
	c.r -= 1
	return c
}

func left(c coord) coord {
	c.c -= 1
	return c
}

func right(c coord) coord {
	c.c += 1
	return c
}

func catch(tail, head coord) coord {
	switch {
	case tail.r == head.r:
		if diff := head.c - tail.c; abs(diff) == 2 {
			tail.c += reduce(diff)
		}
	case tail.c == head.c:
		if diff := head.r - tail.r; abs(diff) == 2 {
			tail.r += reduce(diff)
		}
	default:
		diffc := head.c - tail.c
		diffr := head.r - tail.r
		switch {
		case abs(diffc) == 2 && abs(diffr) == 2:
			tail.c += reduce(diffc)
			tail.r += reduce(diffr)
		case abs(diffc) == 2 && abs(diffr) == 1:
			tail.c += reduce(diffc)
			tail.r += diffr
		case abs(diffc) == 1 && abs(diffr) == 2:
			tail.c += diffc
			tail.r += reduce(diffr)
		}
	}
	return tail
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func reduce(i int) int {
	switch {
	case i < 0:
		return i + 1
	case i > 0:
		return i - 1
	default:
		return i
	}
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func simulate(motions string, knots int) int {
	lines := strings.Split(motions, "\n")

	rope := rope{}
	for i := 0; i < knots; i++ {
		rope.knots = append(rope.knots, coord{})
	}

	visited := set.New(rope.tail())
	for _, line := range lines {
		var fn func(coord) coord

		//println(line)
		fields := strings.Fields(line)
		switch fields[0] {
		case "U":
			fn = up
		case "D":
			fn = down
		case "L":
			fn = left
		case "R":
			fn = right
		}

		steps := mustInt(fields[1])
		for i := 0; i < steps; i++ {
			rope.setHead(fn(rope.head()))

			for k := 1; k < knots; k++ {
				rope.set(k, catch(rope.get(k), rope.get(k-1)))
			}

			visited.Add(rope.tail())
		}
	}
	return visited.Len()
}

// part 1: 6642
// part 2: 2765
func main() {
	fmt.Printf("part 1: %d\n", simulate(input, 2))
	fmt.Printf("part 2: %d\n", simulate(input, 10))
}
