package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mymmrac/x"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(input string) any {
	lines := strings.Split(input, "\n")
	var answer int64

skip:
	for _, line := range lines {
		line = strings.TrimPrefix(line, "Game ")
		gameID, rest, _ := strings.Cut(line, ":")

		c := strings.Split(strings.ReplaceAll(rest[1:], ";", ","), ", ")
		for _, s := range c {
			var count int
			var color string
			_, err := fmt.Sscanf(s, "%d %s", &count, &color)
			x.Assert(err == nil, err)

			if cubes[color] < count {
				continue skip
			}
		}

		gID, err := strconv.ParseInt(gameID, 10, 64)
		x.Assert(err == nil, err)
		answer += gID
	}

	return answer
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	var answer int64

	for _, line := range lines {
		line = strings.TrimPrefix(line, "Game ")
		_, rest, _ := strings.Cut(line, ":")

		c := strings.Split(strings.ReplaceAll(rest[1:], ";", ","), ", ")
		sm := make(map[string]int, 3)
		for _, s := range c {
			var count int
			var color string
			_, err := fmt.Sscanf(s, "%d %s", &count, &color)
			x.Assert(err == nil, err)

			sm[color] = max(count, sm[color])
		}

		var ac int64 = 1
		for _, i := range sm {
			ac *= int64(i)
		}
		answer += ac
	}

	return answer
}
