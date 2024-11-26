package main

import "fmt"

func maxProduct(nums []int) int {
	max, s_max := -1, -1

	for _, num := range nums {
		if num > max {
			s_max = max
			max = num
		} else if num > s_max {
			s_max = num
		}
	}

	return (max - 1) * (s_max - 1)
}

func main() {
	fmt.Println(maxProduct([]int{7, 3}))
}
