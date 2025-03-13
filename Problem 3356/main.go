package main

import "fmt"

func isZeroArray(nums []int, queries [][]int, k int) bool {
	prefix := 0
	diff := make([]int, len(nums) + 1)

    for i:=0; i<k; i++ {
		diff[queries[i][0]]+= queries[i][2] 
		diff[queries[i][1]+1]-= queries[i][2]
	}

	for i:=0; i<len(nums); i++ {
		prefix += diff[i]

		if nums[i] > prefix {
			return false
		}

	}

	return true
}


func minZeroArray(nums []int, queries [][]int) int {

    left, right := 0, len(queries)

	if !isZeroArray(nums, queries, right) {
		return -1
	}

	for left <= right {
		mid := left + (right - left) / 2
		if isZeroArray(nums, queries, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return left
}

func main() {
	fmt.Println(minZeroArray([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}}))
	fmt.Println(minZeroArray([]int{0}, [][]int{{0,0,1},{0,0,1},{0,0,3}}))

}
