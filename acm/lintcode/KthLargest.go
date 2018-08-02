package lintcode

import "sort"

func KthLargest(n int ,nums []int) int {
	sort.Ints(nums)
	return nums[n]
}