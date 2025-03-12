package main

import (
	"fmt"
	"slices"
)

func deckRevealedIncreasing(deck []int) []int {
	slices.Sort(deck)
	queue := make([]int, len(deck))
	result := make([]int, len(deck))

	for i := range queue {
		queue[i] = i
	}

	for _, v := range deck {
		head := queue[0]
		result[head] = v

		if len(queue) > 1 {
			pop := queue[1]
			queue = queue[2:]
			queue = append(queue, pop)
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}
