package lintcode

import (
	"sort"
)

//合并已排序的两个数组
func MergeArray(A, B []int) []int { //调系统库方式
	for i := 0; i < len(A); i++ {
		B = append(B, A[i])
	}
	sort.Ints(B)

	return B
}

func MergeArray2(A, B []int) []int {
	if A == nil || B == nil {
		return nil
	}
	res := make([]int, len(A)+len(B))
	var i, j, index = 0, 0, 0
	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			res[index] = A[i]
			index++
			i++
		} else {
			res[index] = B[j]
			index++
			j++
		}
	}
	for i < len(A) {
		res[index] = A[i]
		index++
		i++
	}
	for j < len(B) {
		res[index] = B[j]
		index++
		j++
	}
	return res
}
