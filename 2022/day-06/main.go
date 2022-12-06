package main

import (
	"fmt"

	"github.com/jwnet/useful/list"
)

func findEndOfMarker(mesg string, markerLen int) int {
	buffer := list.List[rune]{}
	rMap := make(map[rune]int8)

	for i, r := range input {
		buffer.Push(r)
		rMap[r] += 1
		if buffer.Len() == markerLen {
			if len(rMap) == markerLen {
				return i + 1
			}
			end, _ := buffer.Drop()
			if n := rMap[end]; n > 1 {
				rMap[end] = n - 1
			} else {
				delete(rMap, end)
			}
		}
	}
	panic("you must have bad input!")
}

// part 1: 1109
// part 2: 3965
func main() {
	fmt.Printf("part 1: %d\n", findEndOfMarker(input, 4))
	fmt.Printf("part 2: %d\n", findEndOfMarker(input, 14))
}
