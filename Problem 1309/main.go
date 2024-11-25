package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func freqAlphabets(s string) string {

	m1 := regexp.MustCompile(`(1[0-9]|2[0-6])#`)

	s1 := m1.ReplaceAllStringFunc(s, func(s string) string {
		if cc, err := strconv.Atoi(s[0:2]); err == nil {
			return string(cc + 96)
		}
		return ""
	})

	m2 := regexp.MustCompile(`[1-9]`)

	return m2.ReplaceAllStringFunc(s1, func(s string) string {
		return string(byte(s[0]) + 48)
	})

}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}
