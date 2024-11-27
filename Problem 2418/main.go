package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	indexes := make(map[int]int)
	res := make([]string, len(heights))

	for i := 0; i < len(heights); i++ {
		indexes[heights[i]] = i
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})

	for i := 0; i < len(heights); i++ {
		res[i] = names[indexes[heights[i]]]
	}

	return res
}

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
}
