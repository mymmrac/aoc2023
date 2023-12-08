package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   2,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   18727,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   6,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   18024643846273,
			Solution: part2,
		},
	})
}
