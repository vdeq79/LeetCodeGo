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

		for j := g[0] - 1; j >= 0; j-- {
			if matrix[j*n+g[1]] == 0 {
				matrix[j][g[1]] = 1
				total--
			} else if matrix[j][g[1]] < 0 {
				break
			}
		}

		for j := g[0] + 1; j < m; j++ {
			if matrix[j][g[1]] == 0 {
				matrix[j][g[1]] = 1
				total--
			} else if matrix[j][g[1]] < 0 {
				break
			}
		}

		for j := g[1] - 1; j >= 0; j-- {
			if matrix[g[0]][j] == 0 {
				matrix[g[0]][j] = 1
				total--
			} else if matrix[g[0]][j] < 0 {
				break
			}
		}

		for j := g[1] + 1; j < n; j++ {
			if matrix[g[0]][j] == 0 {
				matrix[g[0]][j] = 1
				total--
			} else if matrix[g[0]][j] < 0 {
				break
			}
		}

	}

	return total
}

func main() {
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}
