package main

import "fmt"

func maximumCount(nums []int) int {
    var pos int = -1
	var neg int = len(nums)
	var changed bool = false

	for i, num := range nums {
		if num >= 0 && !changed {
			neg = i
			changed = true
		} 

		if num > 0 {
			pos = len(nums) - i 
			break
		}

	}

	return max(pos, neg)
}

func main() {
	var result = maximumCount([]int{-2,-1,-1,1,2,3})
	fmt.Println(result)
}
