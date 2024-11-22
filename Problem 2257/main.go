package main

import (
	"fmt"
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	var total = m * n
	matrix := make([]byte, m*n)

	for _, w := range walls {
		matrix[w[0]*n+w[1]] = 1
		total--
	}

	for _, g := range guards {
		matrix[g[0]*n+g[1]] = 1
		total--
	}

	for _, g := range guards {
		row := g[0]
		col := g[1]

		for r := row - 1; r >= 0 && matrix[r*n+col] != 1; r-- {
			if matrix[r*n+col] == 0 {
				matrix[r*n+col] = 2
				total--
			}
		}

		for r := row + 1; r < m && matrix[r*n+col] != 1; r++ {
			if matrix[r*n+col] == 0 {
				matrix[r*n+col] = 2
				total--
			}
		}

		for c := col - 1; c >= 0 && matrix[row*n+c] != 1; c-- {
			if matrix[row*n+c] == 0 {
				matrix[row*n+c] = 2
				total--
			}
		}

		for c := col + 1; c < n && matrix[row*n+c] != 1; c++ {
			if matrix[row*n+c] == 0 {
				matrix[row*n+c] = 2
				total--
			}
		}

	}

	return total
}

func main() {
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}
