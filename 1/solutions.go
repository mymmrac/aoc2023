package main

import (
	"strings"
	"unicode"
)

var digits = map[string]int64{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func part1(input string) any {
	lines := strings.Split(input, "\n")
	var answer int64
	for _, line := range lines {
		i := strings.IndexFunc(line, unicode.IsDigit)
		answer += int64(line[i]-'0') * 10
		i = strings.LastIndexFunc(line, unicode.IsDigit)
		answer += int64(line[i] - '0')
	}
	return answer
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	var answer int64
	for _, line := range lines {
		minT := len(line)
		minD := int64(0)
		for text, digit := range digits {
			i := strings.Index(line, text)
			if i >= 0 && i < minT {
				minT = i
				minD = digit
			}
		}
		answer += minD * 10

		maxT := -1
		maxD := int64(0)
		for text, digit := range digits {
			i := strings.LastIndex(line, text)
			if i >= 0 && i > maxT {
				maxT = i
				maxD = digit
			}
		}
		answer += maxD
	}
	return answer
}
