package main

import (
	"fmt"
)

const MOD = 1e9 + 7

func speFibo(n int) int {
	if n == 1 {
		return 2
	} else if n == 2 {
		return 3
	}

	vec := make([]int, n)
	vec[0] = 2
	vec[1] = 3

	for i := 2; i < n; i++ {
		vec[i] = (vec[i-1] + vec[i-2]) % MOD
	}

	return vec[n-1]
}

func countHousePlacements(n int) int {
	s := speFibo(n)
	return (s * s) % MOD
}

func main() {
	fmt.Println(countHousePlacements(1000))
}
