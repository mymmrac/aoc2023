package main

import (
	"strconv"
	"strings"

	"github.com/mymmrac/aoc2023"
)

func part1(input string) any {
	input = aoc2023.TrimSpaces(input)

	times, distances, _ := strings.Cut(input, "\n")
	times = strings.TrimPrefix(times, "Time:")
	distances = strings.TrimPrefix(distances, "Distance:")
	tt := strings.Split(strings.TrimSpace(times), " ")
	dd := strings.Split(strings.TrimSpace(distances), " ")

	var ts []int64
	var ds []int64
	for i := range tt {
		t, _ := strconv.ParseInt(strings.TrimSpace(tt[i]), 10, 64)
		ts = append(ts, t)
		d, _ := strconv.ParseInt(strings.TrimSpace(dd[i]), 10, 64)
		ds = append(ds, d)
	}

	var answer int64 = 1

	for i := 0; i < len(ts); i++ {
		t := ts[i]
		d := ds[i]

		var minT int64
		for j := int64(0); j < t; j++ {
			md := dist(t, j)
			if md > d {
				minT = j
				break
			}
		}
		var maxT int64
		for j := t; j >= 0; j-- {
			md := dist(t, j)
			if md > d {
				maxT = j
				break
			}
		}

		answer *= maxT - minT + 1
	}

	return answer
}

func dist(total, hold int64) int64 {
	return (total - hold) * hold
}

func part2(input string) any {
	input = aoc2023.RemoveSpaces(input)

	times, distances, _ := strings.Cut(input, "\n")
	times = strings.TrimPrefix(times, "Time:")
	distances = strings.TrimPrefix(distances, "Distance:")

	t, _ := strconv.ParseInt(times, 10, 64)
	d, _ := strconv.ParseInt(distances, 10, 64)

	var answer int64 = 1

	var minT int64
	for j := int64(0); j < t; j++ {
		md := dist(t, j)
		if md > d {
			minT = j
			break
		}
	}
	var maxT int64
	for j := t; j >= 0; j-- {
		md := dist(t, j)
		if md > d {
			maxT = j
			break
		}
	}

	answer *= maxT - minT + 1

	return answer
}
