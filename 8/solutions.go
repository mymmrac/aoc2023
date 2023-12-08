package main

import (
	"regexp"
	"strings"

	"github.com/mymmrac/aoc2023"
)

type Move struct {
	L string
	R string
}

func part1(input string) any {
	lines := strings.Split(input, "\n")
	inst := lines[0]
	moves := map[string]Move{}
	rg := regexp.MustCompile(`[(),=]`)

	for _, s := range lines[2:] {
		p := strings.Split(strings.Replace(rg.ReplaceAllString(s, ""), "  ", " ", 1), " ")
		moves[p[0]] = Move{
			L: p[1],
			R: p[2],
		}
	}

	start := "AAA"
	end := "ZZZ"

	var i int64
	curr := start
	for curr != end {
		in := inst[i%int64(len(inst))]

		if in == 'L' {
			curr = moves[curr].L
		} else {
			curr = moves[curr].R
		}

		i++
	}

	return i
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	inst := lines[0]
	moves := map[string]Move{}
	rg := regexp.MustCompile(`[(),=]`)

	var starts, ends []string
	for _, s := range lines[2:] {
		p := strings.Split(strings.Replace(rg.ReplaceAllString(s, ""), "  ", " ", 1), " ")
		moves[p[0]] = Move{
			L: p[1],
			R: p[2],
		}
		if p[0][len(p[0])-1] == 'A' {
			starts = append(starts, p[0])
		}
		if p[0][len(p[0])-1] == 'Z' {
			ends = append(ends, p[0])
		}
	}

	steps := make([]int64, len(starts))
	for k := 0; k < len(starts); k++ {
		var i int64
		curr := starts[k]
		for curr[2] != 'Z' {
			in := inst[i%int64(len(inst))]

			if in == 'L' {
				curr = moves[curr].L
			} else {
				curr = moves[curr].R
			}

			i++
		}
		steps[k] = i
	}

	return aoc2023.LCM(steps[0], steps[1], steps[2:]...)
}
