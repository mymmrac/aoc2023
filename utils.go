package aoc2023

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/mymmrac/x"
	"github.com/stretchr/testify/require"
)

func Read(path string) string {
	data, err := os.ReadFile(path)
	x.Assert(err == nil, path, err)
	return string(data)
}

type Test struct {
	Input    string
	Answer   any
	Solution func(input string) any
}

func RunMany(t *testing.T, tests []Test) {
	for i, test := range tests {
		Run(t, i, test)
	}
}

func Run(t *testing.T, i int, test Test) {
	t.Log("Test:", test.Input)
	answer := test.Solution(Read(test.Input))
	if test.Answer == nil {
		require.FailNow(t, fmt.Sprintf("Answer %d", i+1), answer)
	} else {
		require.EqualValues(t, test.Answer, answer)
	}
}

var cleanSet = regexp.MustCompile(`[ \t]+`)

func TrimSpaces(input string) string {
	input = cleanSet.ReplaceAllString(input, " ")
	input = strings.TrimSpace(input)
	return input
}

func RemoveSpaces(input string) string {
	input = cleanSet.ReplaceAllString(input, "")
	return input
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
