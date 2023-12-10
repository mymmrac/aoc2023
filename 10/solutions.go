package main

import (
	"fmt"
	"slices"
	"strings"
)

type pos struct {
	s  byte
	x  int
	y  int
	dx int
	dy int
}

// 'F', '7', 'J', 'L', '|', '-'

var ppp = map[byte][4][]byte{
	'F': {
		{},              // T
		{'7', 'J', '-'}, // R
		{'J', 'L', '|'}, // B
		{},              // L
	},
	'7': {
		{},              // T
		{},              // R
		{'J', 'L', '|'}, // B
		{'F', 'L', '-'}, // L
	},
	'J': {
		{'F', '7', '|'}, // T
		{},              // R
		{},              // B
		{'F', 'L', '-'}, // L
	},
	'L': {
		{'F', '7', '|'}, // T
		{'7', 'J', '-'}, // R
		{},              // B
		{},              // L
	},
	'|': {
		{'F', '7', '|'}, // T
		{},              // R
		{'J', 'L', '|'}, // B
		{},              // L
	},
	'-': {
		{},              // T
		{'7', 'J', '-'}, // R
		{},              // B
		{'F', 'L', '-'}, // L
	},
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	si := strings.Index(strings.ReplaceAll(input, "\n", ""), "S")
	sp := pos{
		s: 'S',
		x: si % len(grid[0]),
		y: si / len(grid[0]),
	}

	cp := pos{
		x: sp.x,
		y: sp.y,
	}
	t, r, b, l := get(grid, sp.x, sp.y-1), get(grid, sp.x+1, sp.y), get(grid, sp.x, sp.y+1), get(grid, sp.x-1, sp.y)
	for p, c := range ppp {
		if (slices.Contains(c[0], t) || len(c[0]) == 0) &&
			(slices.Contains(c[1], r) || len(c[1]) == 0) &&
			(slices.Contains(c[2], b) || len(c[2]) == 0) &&
			(slices.Contains(c[3], l) || len(c[3]) == 0) {
			cp.s = p
			if slices.Contains([]byte{'F', 'L', '-'}, p) {
				cp.dx = 1
			}
			if slices.Contains([]byte{'7', 'J'}, p) {
				cp.dx = -1
			}
			if slices.Contains([]byte{'|'}, p) {
				cp.dy = 1
			}
			break
		}
	}

	var loop []pos
	loop = append(loop, cp)

	for {
		cp.x += cp.dx
		cp.y += cp.dy
		next := get(grid, cp.x, cp.y)
		if next == 'S' {
			break
		}
		cp.s = next

		nm := ppp[next]
		if len(nm[0]) != 0 && cp.dy == 0 { // T
			cp.dx = 0
			cp.dy = -1
		} else if len(nm[1]) != 0 && cp.dx == 0 { // R
			cp.dx = 1
			cp.dy = 0
		} else if len(nm[2]) != 0 && cp.dy == 0 { // B
			cp.dx = 0
			cp.dy = 1
		} else if len(nm[3]) != 0 && cp.dx == 0 { // L
			cp.dx = -1
			cp.dy = 0
		}

		loop = append(loop, cp)
	}

	return len(loop) / 2
}

func get(grid []string, x, y int) byte {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return 0
	}
	return grid[y][x]
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	si := strings.Index(strings.ReplaceAll(input, "\n", ""), "S")
	sp := pos{
		s: 'S',
		x: si % len(grid[0]),
		y: si / len(grid[0]),
	}

	cp := pos{
		x: sp.x,
		y: sp.y,
	}
	t, r, b, l := get(grid, sp.x, sp.y-1), get(grid, sp.x+1, sp.y), get(grid, sp.x, sp.y+1), get(grid, sp.x-1, sp.y)
	for p, c := range ppp {
		if (slices.Contains(c[0], t) || len(c[0]) == 0) &&
			(slices.Contains(c[1], r) || len(c[1]) == 0) &&
			(slices.Contains(c[2], b) || len(c[2]) == 0) &&
			(slices.Contains(c[3], l) || len(c[3]) == 0) {
			cp.s = p
			sp.s = p
			if slices.Contains([]byte{'F', 'L', '-'}, p) {
				cp.dx = 1
			}
			if slices.Contains([]byte{'7', 'J'}, p) {
				cp.dx = -1
			}
			if slices.Contains([]byte{'|'}, p) {
				cp.dy = 1
			}
			break
		}
	}

	var loop []pos
	loop = append(loop, cp)

	for {
		cp.x += cp.dx
		cp.y += cp.dy
		next := get(grid, cp.x, cp.y)
		if next == 'S' {
			break
		}
		cp.s = next

		ll := loop[len(loop)-1]

		nm := ppp[next]
		if len(nm[0]) != 0 && ll.y != cp.y-1 { // T
			cp.dx = 0
			cp.dy = -1
		} else if len(nm[1]) != 0 && ll.x != cp.x+1 { // R
			cp.dx = 1
			cp.dy = 0
		} else if len(nm[2]) != 0 && ll.y != cp.y+1 { // B
			cp.dx = 0
			cp.dy = 1
		} else if len(nm[3]) != 0 && ll.x != cp.x-1 { // L
			cp.dx = -1
			cp.dy = 0
		} else {
			fmt.Println(string(cp.s), cp.x, cp.y)
			panic("stuck")
		}

		loop = append(loop, cp)
	}

	ss := []byte(grid[sp.y])
	ss[sp.x] = sp.s
	grid[sp.y] = string(ss)

	answer := 0

	for y := 0; y < len(grid); y++ {
		i := 0
		for x := 0; x < len(grid[0]); x++ {
			if inLoop(loop, x, y) {
				switch get(grid, x, y) {
				case 'F':
				innerLoopF:
					for {
						switch get(grid, x, y) {
						case '7':
							break innerLoopF
						case 'J':
							i++
							break innerLoopF
						}
						x++
					}
				case 'L':
				innerLoopL:
					for {
						switch get(grid, x, y) {
						case '7':
							i++
							break innerLoopL
						case 'J':
							break innerLoopL
						}
						x++
					}
				case '|':
					i++
				}
			} else {
				if i%2 == 1 {
					answer++
				}
			}
		}
	}

	return answer
}

func inLoop(loop []pos, x, y int) bool {
	for _, p := range loop {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}
