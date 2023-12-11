package main

import (
	"slices"
	"strings"
)

type pos struct {
	x int
	y int
}

func part1(input string) any {
	grid := strings.Split(input, "\n")

	for y := 0; y < len(grid); y++ {
		if !strings.Contains(grid[y], "#") {
			grid = slices.Insert(grid, y, grid[y])
			y++
		}
	}

hasGalaxy:
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				continue hasGalaxy
			}
		}

		for y := 0; y < len(grid); y++ {
			grid[y] = string(slices.Insert([]byte(grid[y]), x, '.'))
		}
		x++
	}

	galaxies := make(map[pos]struct{})
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				galaxies[pos{
					x: x,
					y: y,
				}] = struct{}{}
			}
		}
	}

	var answer int
	for p1 := range galaxies {
		for p2 := range galaxies {
			if p1 == p2 {
				continue
			}

			d := 0
			poss := map[pos]struct{}{}
			for p1 != p2 {
				dir := pos{}

				switch {
				case p1.y == p2.y && p1.x != p2.x: // Move X
					dir.x = sign(p1.x - p2.x)
				case p1.x == p2.x && p1.y != p2.y: // Move Y
					dir.y = sign(p1.y - p2.y)
				case abs(p1.x-p2.x) > abs(p1.y-p2.y): // Move X
					dir.x = sign(p1.x - p2.x)
				default: // Move Y
					dir.y = sign(p1.y - p2.y)
				}

				p2.x += dir.x
				p2.y += dir.y

				poss[p2] = struct{}{}

				d++
			}

			answer += d
		}
	}

	return answer / 2
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	return 1
}

func part2(input string) any {
	grid := strings.Split(input, "\n")

	expansion := 1000000

	expendedRows := make(map[int]bool)
	expendedCols := make(map[int]bool)

	for y := 0; y < len(grid); y++ {
		if !strings.Contains(grid[y], "#") {
			expendedRows[y] = true
		}
	}

hasGalaxy:
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				continue hasGalaxy
			}
		}

		expendedCols[x] = true
	}

	galaxies := make(map[pos]struct{})
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				galaxies[pos{
					x: x,
					y: y,
				}] = struct{}{}
			}
		}
	}

	var answer int
	for p1 := range galaxies {
		for p2 := range galaxies {
			if p1 == p2 {
				continue
			}

			d := 0
			for p1 != p2 {
				dir := pos{}

				switch {
				case p1.y == p2.y && p1.x != p2.x: // Move X
					dir.x = sign(p1.x - p2.x)
				case p1.x == p2.x && p1.y != p2.y: // Move Y
					dir.y = sign(p1.y - p2.y)
				case abs(p1.x-p2.x) > abs(p1.y-p2.y): // Move X
					dir.x = sign(p1.x - p2.x)
				default: // Move Y
					dir.y = sign(p1.y - p2.y)
				}

				p2.x += dir.x
				p2.y += dir.y

				m := 1
				if dir.x != 0 && expendedCols[p2.x] || dir.y != 0 && expendedRows[p2.y] {
					m = expansion
				}
				d += m
			}

			answer += d
		}
	}

	return answer / 2
}
