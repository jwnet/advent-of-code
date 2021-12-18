package main

import (
	"fmt"
	"math"
)

type target struct {
	xmin, xmax int
	ymin, ymax int
}

type probe struct {
	x, y   int
	vx, vy int
}

func (p *probe) step() {
	p.x += p.vx
	p.y += p.vy
	switch {
	case p.vx > 0:
		p.vx--
	case p.vx < 0:
		p.vx++
	}
	p.vy--
}

func (p *probe) inRange(t target) bool {
	return t.xmin <= p.x && p.x <= t.xmax && // in x range
		t.ymin <= p.y && p.y <= t.ymax // in y range

}

func (p *probe) beyond(t target) bool {
	return p.y < t.ymin || p.x > t.xmax
}

func vxmin(t target) int {
	var vxmin int
	switch {
	case t.xmin < 0:
		vxmin = t.xmin
	case t.xmin > 0:
		// CeiL(x), solving for x in x(x+1)/2 = t.xmin,
		// which is the lowest possible velocity for x.
		vxmin = int(math.Ceil(math.Sqrt(8*float64(t.xmin))/2 - 0.5))
	default:
		vxmin = 0
	}
	return vxmin
}

func vymax(t target) int {
	if t.ymin < 0 && t.ymax < 0 {
		return (-1 * t.ymin) - 1
	}		
	panic("what kind of \"trench\" is this")
}

func part1(t target) int {
	vymax := vymax(t)
	return vymax * (vymax + 1) / 2
}

func part2(t target) int {
	vxmin := vxmin(t)
	vymax := vymax(t)
	count := 0
	for vy := t.ymin; vy <= vymax; vy++ {
		for vx := vxmin; vx <= t.xmax; vx++ {
			p := probe{x: 0, y: 0, vx: vx, vy: vy}
			for !p.beyond(t) {
				p.step()
				if p.inRange(t) {
					count++
					break
				}
			}
		}
	}
	return count
}

func main() {
	sample := target{xmin: 20, xmax: 30, ymin: -10, ymax: -5}
	input := target{xmin: 88, xmax: 125, ymin: -157, ymax: -103}

	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\t input: %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(sample))
	fmt.Printf("\t input: %d\n", part2(input))

	// Part One
	// 	 sample: 45
	// 	 input: 12246
	// Part Two
	// 	sample: 112
	// 	 input: 3528
}
