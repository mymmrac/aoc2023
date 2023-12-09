package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/mymmrac/x"
)

func part1(input string) any {
	lines := strings.Split(input, "\n")
	var answer int64
	for _, line := range lines {
		seq := x.Convert(strings.Split(line, " "), func(item string) int64 {
			i, err := strconv.ParseInt(item, 10, 64)
			x.Assert(err == nil, err)
			return i
		})

		var diffs [][]int64
		diffs = append(diffs, seq)
		diffs = append(diffs, diffSeq(seq))
		for !isConst(diffs[len(diffs)-1]) {
			diffs = append(diffs, diffSeq(diffs[len(diffs)-1]))
		}

		ld := diffs[len(diffs)-1]
		diffs[len(diffs)-1] = append(diffs[len(diffs)-1], ld[len(ld)-1]+ld[1]-ld[0])

		for i := len(diffs) - 1; i >= 1; i-- {
			nd := diffs[i-1]
			d := diffs[i][len(diffs[i])-1]
			diffs[i-1] = append(nd, nd[len(nd)-1]+d)
		}

		ldd := diffs[0]
		answer += ldd[len(ldd)-1]
	}
	return answer
}

func diffSeq(seq []int64) []int64 {
	ns := make([]int64, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		ns[i] = seq[i+1] - seq[i]
	}
	return ns
}

func isConst(seq []int64) bool {
	d := seq[1] - seq[0]
	for i := 0; i < len(seq)-1; i++ {
		if seq[i+1]-seq[i] != d {
			return false
		}
	}
	return true
}

func part2(input string) any {
	lines := strings.Split(input, "\n")
	var answer int64
	for _, line := range lines {
		seq := x.Convert(strings.Split(line, " "), func(item string) int64 {
			i, err := strconv.ParseInt(item, 10, 64)
			x.Assert(err == nil, err)
			return i
		})

		var diffs [][]int64
		diffs = append(diffs, seq)
		diffs = append(diffs, diffSeq(seq))
		for !isConst(diffs[len(diffs)-1]) {
			diffs = append(diffs, diffSeq(diffs[len(diffs)-1]))
		}

		ld := diffs[len(diffs)-1]
		diffs[len(diffs)-1] = slices.Insert(diffs[len(diffs)-1], 0, ld[0]-(ld[1]-ld[0]))

		for i := len(diffs) - 1; i >= 1; i-- {
			nd := diffs[i-1]
			d := diffs[i][0]
			diffs[i-1] = slices.Insert(nd, 0, nd[0]-d)
		}

		// for i, diff := range diffs {
		// 	fmt.Println(strings.Repeat(" ", i), diff)
		// }
		// fmt.Println()

		answer += diffs[0][0]
	}
	return answer
}
