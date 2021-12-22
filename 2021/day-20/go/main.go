package main

import (
	"fmt"
	"strings"
)

type infbit int

const (
	dot  infbit = 0
	hash infbit = 511
)

type imager struct {
	algo      string
	image     []string
	inf       byte
	nenhances int
}

func (img *imager) initInf() {
	img.inf = img.algo[dot]
}

func (img *imager) enhance() {
	img.nenhances += 1
	if img.inf == 0 {
		img.initInf()
	}
	switch img.inf {
	case '.':
		img.inf = img.algo[dot]
	case '#':
		img.inf = img.algo[hash]
	}
	newImage := make([]string, len(img.image)+2)
	for i := -1; i < len(img.image)+1; i++ {
		newStr := strings.Builder{}
		for j := -1; j < len(img.image[0])+1; j++ {
			newStr.WriteByte(img.lookupEnhance(i, j))
		}
		newImage[i+1] = newStr.String()
	}
	img.image = newImage
}

func (img *imager) lookupEnhance(i, j int) byte {
	idx := 0
	for ii := i - 1; ii <= i+1; ii++ {
		for jj := j - 1; jj <= j+1; jj++ {
			idx = (idx << 1) + img.getBit(ii, jj)
		}
	}
	return img.algo[idx]
}

func (img *imager) getBit(i, j int) int {
	if i < 0 || i >= len(img.image) || j < 0 || j >= len(img.image[0]) {
		switch img.inf {
		case '.':
			return 0
		case '#':
			return 1
		}
	}
	if b := img.image[i][j]; rune(b) == '.' {
		return 0
	}
	return 1
}

func (i imager) String() string {
	sb := strings.Builder{}
	for _, str := range i.image {
		sb.WriteString(str)
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (img *imager) countLit() int {
	count := 0
	for _, str := range img.image {
		for _, r := range str {
			if r == '#' {
				count++
			}
		}
	}
	return count
}

func part1(img *imager) int {
	img.enhance()
	img.enhance()
	return img.countLit()
}

func part2(img *imager) int {
	for i := img.nenhances; i < 50; i++ {
		img.enhance()
	}
	return img.countLit()
}

func main() {
	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(&sample))
	fmt.Printf("\t input: %d\n", part1(&input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(&sample))
	fmt.Printf("\t input: %d\n", part2(&input))

	// Part One
	// 	sample: 35
	// 	 input: 5379
	// Part Two
	// 	sample: 3351
	// 	 input: 17917
}
