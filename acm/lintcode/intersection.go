package lintcode

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int { //lintcode 547
	// write your code here
	if len(nums1) == 0 || len(nums2) == 0 {
		fmt.Println("asd")
		return []int{}
	}
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
	sort.Ints(ret)
	return ret
}
func intersection2(nums1 []int, nums2 []int) []int { //lintcode 548
	// write your code here
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
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
	fmt.Println(m)
	for _, value := range nums2 {
		v, ok := m[value]
		if ok {
			m[value] = v - 1
		}
	}
	fmt.Println(m)
	for key, value := range m {
		for i := 0; i < value; i++ {
			ret = append(ret, key)
		}
	}
	sort.Ints(ret)
	return ret
}
