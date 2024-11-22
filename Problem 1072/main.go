package main

import (
	"bytes"
	"fmt"
)

func maxEqualRowsAfterFlips(matrix [][]int) int {
	dict := make(map[string]int)
	res := 0

	for _, r := range matrix {
		var s bytes.Buffer
		for _, c := range r {
			if c == r[0] {
				s.WriteString("0")
			} else {
				s.WriteString("1")
			}
		}

		dict[s.String()]++
		res = max(res, dict[s.String()])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxEqualRowsAfterFlips([][]int{{0, 1}, {1, 0}}))
}
