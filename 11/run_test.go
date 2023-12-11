package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   374,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   9947476,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   82000210,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   519939907614,
			Solution: part2,
		},
	})
}
