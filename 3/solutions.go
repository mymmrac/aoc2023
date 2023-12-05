package main

import (
	"strconv"
	"strings"
	"unicode"
)

func part1(input string) any {
	grid := strings.Split(input, "\n")
	parts := map[int]map[int]bool{}
	for y, line := range grid {
		for x, s := range line {
			if s == '.' {
				continue
			}
			if unicode.IsDigit(s) {
				continue
			}
			for i := -1; i <= 1; i++ {
				yy := y + i
				if yy < 0 || yy >= len(grid) {
					continue
				}
				for j := -1; j <= 1; j++ {
					xx := x + j
					if xx < 0 || xx >= len(grid[0]) {
						continue
					}
					if _, ok := parts[yy]; !ok {
						parts[yy] = map[int]bool{}
					}
					parts[yy][xx] = true
				}
			}
		}
	}

	var answer int64
	for y, line := range grid {
		x := 0
		for x < len(line) {
			s := rune(line[x])

			if unicode.IsDigit(s) {
				number := ""
				ok := false

				for unicode.IsDigit(s) && x < len(line) {
					number += string(s)

					if parts[y][x] {
						ok = true
					}

					x++
					if x < len(line) {
						s = rune(line[x])
					}
				}

				if ok {
					a, _ := strconv.ParseInt(number, 10, 64)
					answer += a
				}
			}
			x++
		}
	}

	// for y, line := range grid {
	// 	for x, _ := range line {
	// 		if parts[y][x] {
	// 			fmt.Print("X")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return answer
}

func part2(input string) any {
	grid := strings.Split(input, "\n")

	type point struct {
		r rune
		x int
		y int
	}
	parts := map[int]map[int]point{}

	for y, line := range grid {
		for x, s := range line {
			if s == '.' {
				continue
			}
			if unicode.IsDigit(s) {
				continue
			}

			for i := -1; i <= 1; i++ {
				yy := y + i
				if yy < 0 || yy >= len(grid) {
					continue
				}
				for j := -1; j <= 1; j++ {
					xx := x + j
					if xx < 0 || xx >= len(grid[0]) {
						continue
					}
					if _, ok := parts[yy]; !ok {
						parts[yy] = map[int]point{}
					}
					parts[yy][xx] = point{
						r: s,
						x: x,
						y: y,
					}
				}
			}
		}
	}

	ratios := map[point][]int64{}

	for y, line := range grid {
		x := 0
		for x < len(line) {
			s := rune(line[x])

			if unicode.IsDigit(s) {
				number := ""
				ok := false

				var p point
				for unicode.IsDigit(s) && x < len(line) {
					number += string(s)

					if pp, found := parts[y][x]; found {
						p = pp
						ok = true
					}

					x++
					if x < len(line) {
						s = rune(line[x])
					}
				}

				if ok && p.r == '*' {
					a, _ := strconv.ParseInt(number, 10, 64)
					ratios[p] = append(ratios[p], a)
				}
			}
			x++
		}
	}

	var answer int64
	for _, rr := range ratios {
		if len(rr) == 2 {
			answer += rr[0] * rr[1]
		}
	}

	return answer
}
