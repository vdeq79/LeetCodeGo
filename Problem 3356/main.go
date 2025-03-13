package main

import "fmt"

func minZeroArray(nums []int, queries [][]int) int {
	sum := 0

	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}

	if sum == 0 {
		return 0
	}

    for j, query := range queries {
		for i:=query[0]; i<=query[1]; i++ {

			if(nums[i] > 0){
				substract := min(nums[i], query[2])
				nums[i] -= substract
				sum -= substract

				if sum == 0 {
					return j+1
				}
			}

		}
	}

	return -1
}

func main() {
	fmt.Println(minZeroArray([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}}))
}
