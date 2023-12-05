package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   4361,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   540025,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   467835,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   84584891,
			Solution: part2,
		},
	})
}
