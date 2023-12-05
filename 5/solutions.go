package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/mymmrac/x"
)

func part1(input string) any {
	lines := strings.Split(input, "\n")
	seedsStr := strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ")
	_ = seedsStr

	var seeds []uint64
	for _, s := range seedsStr {
		sc, err := strconv.ParseUint(s, 10, 64)
		x.Assert(err == nil, err)
		seeds = append(seeds, sc)
	}

	to := ""
	mapping := map[string][]Mapping{}

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			_, rest, _ := strings.Cut(line, "-")
			to, _, _ = strings.Cut(rest[3:], " ")
			mapping[to] = []Mapping{}
			continue
		}
		m := Mapping{}
		_, err := fmt.Sscanf(line, "%d %d %d", &m.Destination, &m.Source, &m.Count)
		x.Assert(err == nil, err)
		mapping[to] = append(mapping[to], m)
	}

	answer := uint64(math.MaxUint64)
	for _, seed := range seeds {
		curr := seed

		to = "soil"
	loop:
		for {
			mm := mapping[to]
			curr = find(mm, curr)

			switch to {
			case "soil":
				to = "fertilizer"
			case "fertilizer":
				to = "water"
			case "water":
				to = "light"
			case "light":
				to = "temperature"
			case "temperature":
				to = "humidity"
			case "humidity":
				to = "location"
			case "location":
				break loop
			}
		}

		answer = min(curr, answer)
	}

	return answer
}

func find(mm []Mapping, source uint64) uint64 {
	for _, m := range mm {
		if m.Source <= source && source <= m.Source+m.Count {
			return m.Destination + (source - m.Source)
		}
	}
	return source
}

type Mapping struct {
	Source      uint64
	Destination uint64
	Count       uint64
}

// TODO: Optimize this solution (very slow)
func part2(input string) any {
	lines := strings.Split(input, "\n")
	seedsStr := strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ")
	_ = seedsStr

	to := ""
	mapping := map[string][]Mapping{}

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if !unicode.IsDigit(rune(line[0])) {
			_, rest, _ := strings.Cut(line, "-")
			to, _, _ = strings.Cut(rest[3:], " ")
			mapping[to] = []Mapping{}
			continue
		}
		m := Mapping{}
		_, err := fmt.Sscanf(line, "%d %d %d", &m.Destination, &m.Source, &m.Count)
		x.Assert(err == nil, err)
		mapping[to] = append(mapping[to], m)
	}

	answer := uint64(math.MaxUint64)

	var seeds []uint64
	for i := 0; i < len(seedsStr); i += 2 {
		{
			seed, err := strconv.ParseUint(seedsStr[i], 10, 64)
			x.Assert(err == nil, err)

			count, err := strconv.ParseUint(seedsStr[i+1], 10, 64)
			x.Assert(err == nil, err)

			for j := seed; j < seed+count; j++ {
				seeds = append(seeds, j)
			}
		}

		for _, seed := range seeds {
			curr := seed
			to = "soil"
		loop:
			for {
				mm := mapping[to]
				curr = find(mm, curr)

				switch to {
				case "soil":
					to = "fertilizer"
				case "fertilizer":
					to = "water"
				case "water":
					to = "light"
				case "light":
					to = "temperature"
				case "temperature":
					to = "humidity"
				case "humidity":
					to = "location"
				case "location":
					break loop
				}
			}

			answer = min(curr, answer)
		}

		seeds = nil
		fmt.Println((i/2)+1, answer)
	}

	return answer
}
