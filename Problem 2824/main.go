package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, target int, left int, right int) int {

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left

}

func countPairs(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	count := 0

	for i := 0; i < len(nums); i++ {

		index := binarySearch(nums, target-nums[i]-1, i+1, len(nums)-1)

		if index < len(nums) {
			count += index - i

			if nums[index] > target-nums[i]-1 {
				count--
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countPairs([]int{-5, 0, -7, -1, 9, 8, -9, 9}, -14))
}
