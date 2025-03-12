package main

import "fmt"

func sortColors(nums []int)  {
	left, right:= 0, len(nums)-1
    
	for left < len(nums) && nums[left] == 0{
		left++
	}

	for right >= 0 && nums[right] == 2{
		right-- 
	}

	for i:=left; i<=right; i++{

		if nums[i] == 0{
			nums[i] = nums[left]
			nums[left] = 0
			for left < len(nums) && nums[left] == 0{
				left++
			}
			i = left - 1
		}else if nums[i] == 2{
			nums[i] = nums[right]
			nums[right] = 2
			for right >= 0 && nums[right] == 2{
				right-- 
			}

			i--
		}

	}
}

func main() {
	var nums = []int{1,2,0}
	sortColors(nums)
	fmt.Println(nums);
}
