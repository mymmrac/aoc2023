package main

import "strings"

func part1(input string) any {
	lines := strings.Split(input, "\n")
	var answer int
	for _, line := range lines {
		_, cards, _ := strings.Cut(line, ":")
		winning, my, _ := strings.Cut(cards, "|")

		w := map[string]bool{}
		for _, s := range strings.Split(strings.TrimSpace(winning), " ") {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			w[s] = true
		}

		var points int
		for _, s := range strings.Split(strings.TrimSpace(my), " ") {
			if w[strings.TrimSpace(s)] {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		answer += points
	}
	return answer
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	cardsCount := map[int]int{}
	for i, line := range lines {
		_, cards, _ := strings.Cut(line, ":")
		winning, my, _ := strings.Cut(cards, "|")

		w := map[string]bool{}
		for _, s := range strings.Split(strings.TrimSpace(winning), " ") {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			w[s] = true
		}

		var count int
		for _, s := range strings.Split(strings.TrimSpace(my), " ") {
			if w[strings.TrimSpace(s)] {
				count++
			}
		}

		cardsCount[i+1] += 1
		for j := (i + 1) + 1; j <= (i+1)+count; j++ {
			cardsCount[j] += cardsCount[i+1]
		}
	}

	var answer int
	for _, c := range cardsCount {
		answer += c
	}

	return answer
}
