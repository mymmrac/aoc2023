package main

import (
	"testing"

	"github.com/mymmrac/aoc2023"
)

func TestRun(t *testing.T) {
	aoc2023.RunMany(t, []aoc2023.Test{
		{
			Input:    "input0.txt",
			Answer:   142,
			Solution: part1,
		},
		{
			Input:    "input1.txt",
			Answer:   55538,
			Solution: part1,
		},
		{
			Input:    "input2.txt",
			Answer:   281,
			Solution: part2,
		},
		{
			Input:    "input1.txt",
			Answer:   54875,
			Solution: part2,
		},
	})
}
