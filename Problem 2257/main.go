package main

import (
	"fmt"
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	var total = m * n
	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := range walls {
		matrix[walls[i][0]][walls[i][1]] = -1
		total--
	}

	for i := range guards {
		matrix[guards[i][0]][guards[i][1]] = -2
		total--
	}

	for i := range guards {

		for j := guards[i][0] - 1; j >= 0; j-- {
			if matrix[j][guards[i][1]] == 0 {
				matrix[j][guards[i][1]] = 1
				total--
			} else if matrix[j][guards[i][1]] < 0 {
				break
			}
		}

		for j := guards[i][0] + 1; j < m; j++ {
			if matrix[j][guards[i][1]] == 0 {
				matrix[j][guards[i][1]] = 1
				total--
			} else if matrix[j][guards[i][1]] < 0 {
				break
			}
		}

		for j := guards[i][1] - 1; j >= 0; j-- {
			if matrix[guards[i][0]][j] == 0 {
				matrix[guards[i][0]][j] = 1
				total--
			} else if matrix[guards[i][0]][j] < 0 {
				break
			}
		}

		for j := guards[i][1] + 1; j < n; j++ {
			if matrix[guards[i][0]][j] == 0 {
				matrix[guards[i][0]][j] = 1
				total--
			} else if matrix[guards[i][0]][j] < 0 {
				break
			}
		}

	}

	return total
}

func main() {
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
}
