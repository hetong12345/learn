package lintcode

import (
	"fmt"
)

func OdigitCounts(k int, n int) int { //lintcode第3题，统计数字
	res := 0
	for i := k; i <= n; i++ {
		res += single(i, k)
	}
	return res
}

func single(i, k int) int {
	fmt.Println("asd")
	if i == 0 && k == 0 {
		return 1
	}
	var si int
	for {
		if i%10 == k {
			si++
		}
		i = i / 10

		if i == 0 {
			break
		}
	}

	return si
}
