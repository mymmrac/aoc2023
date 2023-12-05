package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   nil,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   nil,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   nil,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   nil,
			Solution: part2,
		},
	})
}
