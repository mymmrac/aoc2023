package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   8,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   2331,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   2286,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   71585,
			Solution: part2,
		},
	})
}
