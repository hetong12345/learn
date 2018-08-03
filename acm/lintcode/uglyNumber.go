package lintcode

import (
	"math"
)

func UglyNumber(num int) int { //寻找第n个丑数
	var ugly []float64
	var p2, p3, p5 int
	ugly = append(ugly, 1.0)
	for i := 0; i < num+3; i++ {
		next := math.Min(math.Min(ugly[p2]*2, ugly[p3]*3), ugly[p5]*5)
		ugly = append(ugly, next)
		for next >= ugly[p2]*2 {
			p2++
		}
		for next >= ugly[p3]*3 {
			p3++
		}
		for next >= ugly[p5]*5 {
			p5++
		}
	}
	return int(ugly[num-1])
}
func RemoveDuplicatesAndEmpty(a []int) (ret []int) {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
