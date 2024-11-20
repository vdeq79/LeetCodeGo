package main

import "fmt"

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		val, ok := table[complement]
		if ok {
			return []int{val, i}
		} else {
			table[nums[i]] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
