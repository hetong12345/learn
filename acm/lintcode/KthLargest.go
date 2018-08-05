package lintcode

import "sort"

func KthLargest(n int, nums []int) int {//第K大的元素
	sort.Ints(nums)
	return nums[len(nums)-n]
}
