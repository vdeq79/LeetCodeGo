package main

import "fmt"

func minZeroArray(nums []int, queries [][]int) int {

	diff := make([]int, len(nums) + 1)
	prefix := 0
	k := 0

    for i:=0; i<len(nums); i++ {
		prefix += diff[i]

		for nums[i] > prefix && k < len(queries) {

			if queries[k][1] >= i{

				if (queries[k][0] <= i){
					diff[i] += queries[k][2]
					prefix += queries[k][2]
				} else {
					diff[queries[k][0]] += queries[k][2]
				}

				diff[queries[k][1] + 1] -= queries[k][2]
			}

			k++
		}

		if nums[i] > prefix {
			return -1
		}
	}

	return k
}

func main() {
	fmt.Println(minZeroArray([]int{0,10}, [][]int{{0,1,2},{0,0,2},{0,1,2},{1,1,4},{0,1,3},{1,1,4},{0,1,2},{0,1,2},{0,1,2},{0,0,2},{1,1,2},{0,0,2},{0,0,3},{1,1,3},{0,0,5}}))

}
