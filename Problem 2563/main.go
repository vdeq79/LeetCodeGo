package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, target int, left int, right int) int {

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left

}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var count int64 = 0

	for i := 0; i < len(nums); i++ {

		leftIndex := binarySearch(nums, lower-nums[i], i+1, len(nums)-1)
		rightIndex := binarySearch(nums, upper-nums[i]+1, i+1, len(nums)-1)

		count += int64(rightIndex - leftIndex)
	}

	return count
}

func main() {
	fmt.Println(countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
}
