package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   114,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   1842168671,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   5,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   903,
			Solution: part2,
		},
	})
}
