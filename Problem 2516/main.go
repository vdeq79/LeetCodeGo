package main

import (
	"fmt"
)

func takeCharacters(s string, k int) int {
	// Total counts
	count := make([]int, 3)
	for _, c := range s {
		count[c-'a']++
	}

	if min(count[0], min(count[1], count[2])) < k {
		return -1
	}

	res := len(s)
	r := 0

	for l := 0; l < len(s); l++ {
		count[s[l]-'a']--

		for min(count[0], min(count[1], count[2])) < k {
			count[s[r]-'a']++
			r++
		}

		res = min(res, len(s)-l+r-1)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
}
