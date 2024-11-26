package main

import "fmt"

func findChampion(n int, edges [][]int) int {

	champArray := make([]byte, n)

	for _, edge := range edges {
		champArray[edge[1]]++
	}

	index := -1

	for i := 0; i < n; i++ {
		if champArray[i] == 0 {
			if index == -1 {
				index = i
			} else {
				return -1
			}

		}
	}

	return index
}

func main() {
	fmt.Println(findChampion(3, [][]int{{0, 1}, {1, 2}}))
}
