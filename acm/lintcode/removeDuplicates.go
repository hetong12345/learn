package lintcode

import (
	"fmt"
)

func RemoveDuplicates(nums []int) int { //移除重复数
	if len(nums) == 0 {
		return 0
	} else {
		var j = 1
		for i := 1; i < len(nums); i++ {
			if nums[i] == nums[j-1] {

			} else {
				j++
			}
			fmt.Println(nums[i])
		}
		return 1
	}
}
