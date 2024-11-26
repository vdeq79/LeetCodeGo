package main

import (
	"fmt"
)

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func maxVowels(s string, k int) int {
	cur, max := 0, 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			cur++
		}
	}

	max = cur

	for i := 1; i <= len(s)-k; i++ {
		if isVowel(s[i-1]) {
			cur--
		}

		if isVowel(s[i+k-1]) {
			cur++
		}

		if cur > max {
			max = cur
		}
	}

	return max
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}
