package lintcode

import (
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int { //lintcode 547
	m := map[int]int{}
	var ret []int
	for _, value := range nums1 {
		m[value] = 1
	}
	for _, value := range nums2 {
		if m[value] == 1 {
			ret = append(ret, value)
			m[value] = 0
		}
	}
	if len(ret) == 0 {
		return []int{}
	}
	sort.Ints(ret)
	return ret
}
func intersection2(nums1 []int, nums2 []int) []int { //lintcode 548
	m := map[int]int{}
	var ret []int
	for _, value := range nums1 {
		v, ok := m[value]
		if ok {
			m[value] = v + 1
		} else {
			m[value] = 1
		}
	}
	for _, value := range nums2 {
		v := m[value]
		if v > 0 {
			m[value] = v - 1
			ret = append(ret, value)
		}
	}
	if len(ret) == 0 {
		return []int{}
	}
	sort.Ints(ret)
	return ret
}
