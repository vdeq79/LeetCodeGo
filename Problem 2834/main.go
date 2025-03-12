package main

import "fmt"

func minimumPossibleSum(n int, target int) int {
	var half int = target / 2

	fmt.Println(half)

	if n < half {
		return (n * (n + 1) / 2) % (1000000007)
	}

	return ((n-half)*target + (n-half)*(n-half-1)/2 + (half)*(half+1)/2) % (1000000007)

}

func main() {
	fmt.Println(minimumPossibleSum(39636, 49035))
}
