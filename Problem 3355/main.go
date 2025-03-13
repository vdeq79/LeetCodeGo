package main

import "fmt"

func isZeroArray(nums []int, queries [][]int) bool {
	diff := make([]int, len(nums) + 1)
	prefix := 0

    for _, query := range queries {
		diff[query[0]]++ 
		diff[query[1]+1]-- 	
	}

	for i:=0; i<len(nums); i++ {
		prefix += diff[i]

		if nums[i] > prefix {
			return false
		}

	}

	return true
}

func main() {
	fmt.Println(isZeroArray([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}}))
}
