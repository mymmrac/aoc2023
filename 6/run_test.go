package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   288,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   500346,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   71503,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   42515755,
			Solution: part2,
		},
	})
}
