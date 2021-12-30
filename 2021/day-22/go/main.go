package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type prange struct {
	p1, p2 int
}

func (p prange) abs() int {
	return int(math.Abs(float64(p.p2 - p.p1)))
}

func (r prange) cut(o prange) (remainder []prange, cut prange) {
	if o.p1 <= r.p1 && r.p2 <= o.p2 {
		return []prange{}, r
	}
	if o.p1 <= r.p1 && o.p2 < r.p2 {
		return []prange{{o.p2 + 1, r.p2}}, prange{r.p1, o.p2}
	}
	if r.p1 < o.p1 && r.p2 <= o.p2 {
		return []prange{{r.p1, o.p1 - 1}}, prange{o.p1, r.p2}
	}
	return []prange{{r.p1, o.p1 - 1}, {o.p2 + 1, r.p2}}, o
}

func (r prange) intersects(o prange) bool {
	return (r.p1 <= o.p1 && o.p1 <= r.p2) ||
		(r.p1 <= o.p2 && o.p2 <= r.p2) ||
		(o.p1 <= r.p1 && r.p1 <= o.p2) ||
		(o.p1 <= r.p2 && r.p2 <= o.p2)
}

type sector struct {
	x, y, z prange
}

func (s sector) intersects(o sector) bool {
	return s.x.intersects(o.x) &&
		s.y.intersects(o.y) &&
		s.z.intersects(o.z)
}

func (s sector) contains(o sector) bool {
	return (s.x.p1 <= o.x.p1 && o.x.p2 <= s.x.p2) &&
		(s.y.p1 <= o.y.p1 && o.y.p2 <= s.y.p2) &&
		(s.z.p1 <= o.z.p1 && o.z.p2 <= s.z.p2)
}

func (s sector) ncubes() int {
	return (s.x.abs() + 1) * (s.y.abs() + 1) * (s.z.abs() + 1)
}

func (s sector) cut(o sector) (remainder []sector) {
	xs, x := s.x.cut(o.x)
	for _, r := range xs {
		remainder = append(remainder, sector{r, s.y, s.z})
	}
	ys, y := s.y.cut(o.y)
	for _, r := range ys {
		remainder = append(remainder, sector{x, r, s.z})
	}
	zs, _ := s.z.cut(o.z)
	for _, r := range zs {
		remainder = append(remainder, sector{x, y, r})
	}
	return
}

type instruction struct {
	on bool
	s  sector
}

func toInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

func minmax(x, y int) (int, int) {
	if x <= y {
		return int(x), int(y)
	}
	return int(y), int(x)
}

func readInstructions(file string) []instruction {
	instructions := []instruction{}
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Could not open file %s: %v", file, err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ins := instruction{}
		line := scanner.Text()
		if line[:2] == "on" {
			ins.on = true
			line = line[3:]
		} else {
			line = line[4:]
		}
		ranges := strings.Split(line, ",")
		for i, r := range ranges {
			ranges[i] = r[2:]
		}
		s := sector{}
		xr := strings.Split(ranges[0], "..")
		s.x.p1, s.x.p2 = minmax(toInt(xr[0]), toInt(xr[1]))
		yr := strings.Split(ranges[1], "..")
		s.y.p1, s.y.p2 = minmax(toInt(yr[0]), toInt(yr[1]))
		zr := strings.Split(ranges[2], "..")
		s.z.p1, s.z.p2 = minmax(toInt(zr[0]), toInt(zr[1]))
		ins.s = s
		instructions = append(instructions, ins)
	}
	return instructions
}

func reboot(instructions []instruction, initOnly bool) int {
	sectors := []sector{}
	initSector := sector{prange{-50, 50}, prange{-50, 50}, prange{-50, 50}}
	for _, ins := range instructions {
		if initOnly && !initSector.contains(ins.s) {
			continue
		}
		if len(sectors) == 0 && ins.on {
			sectors = append(sectors, ins.s)
			continue
		}
		switch ins.on {
		case true: // turn on cubes
			rs := []sector{ins.s}
			for _, s := range sectors {
				nrs := []sector{}
				for _, r := range rs {
					if r.intersects(s) {
						nrs = append(nrs, r.cut(s)...)
					} else {
						nrs = append(nrs, r)
					}
				}
				rs = nrs
			}
			sectors = append(sectors, rs...)

		case false: // turn off cubes
			ns := make([]sector, 0, len(sectors))
			for _, s := range sectors {
				if s.intersects(ins.s) {
					ns = append(ns, s.cut(ins.s)...)
				} else {
					ns = append(ns, s)
				}
			}
			sectors = ns
		}
	}

	ncubes := 0
	for _, s := range sectors {
		ncubes += s.ncubes()
	}
	return ncubes
}

func part1(instructions []instruction) int {
	return reboot(instructions, true)
}

func part2(instructions []instruction) int {
	return reboot(instructions, false)
}

func main() {
	sample1 := readInstructions("../data/sample1")
	sample2 := readInstructions("../data/sample2")
	input := readInstructions("../data/input")

	fmt.Println("Part One")
	fmt.Printf("\tsample1: %d\n", part1(sample1))
	fmt.Printf("\t  input: %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample2: %d\n", part2(sample2))
	fmt.Printf("\t  input: %d\n", part2(input))

	// Part One
	// 	sample1: 590784
	// 	  input: 606484
	// Part Two
	// 	sample2: 2758514936282235
	// 	  input: 1162571910364852
}
