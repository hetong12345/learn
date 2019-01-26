package algorithm

import (
	"math/rand"
	"time"
)

func CreateRandomArray(len, lRange, rRange int) []int { //创建随机数组
	if lRange > rRange {
		panic("range error")
	}
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, len)
	for key := range arr {
		arr[key] = rand.Intn(rRange-lRange+1) + lRange
	}
	return arr
}

func TestSort(sort func([]int) []int, arr []int) time.Duration { //测试排序算法
	starTime := time.Now()
	sort(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			panic("sort fail")
		}
	}
	return time.Since(starTime)
}

func CreateDeduplicationArray(len, lRange, rRange int) []int { //创建去重数组
	if lRange > rRange {
		panic("range error")
	}
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, len)
	for key := range arr {
		arr[key] = rand.Intn(rRange-lRange+1) + lRange
	}
	return arr
}
