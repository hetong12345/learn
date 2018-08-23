package lintcode

import (
	// "fmt"
	"math"
	// "strconv"
)

func Ofactorial(num int64) int64 { //计算阶乘0的个数
	// fmt.Println(num / 5.0)
	var tmp = num
	var n int
	var res int64
	var i = 1
	for {
		tmp = tmp / 5.0
		if tmp < 1 {
			break
		}
		n++
	}
	// fmt.Println(tmp, n)
	for ; i <= n; i++ {
		res += int64(float64(num) / math.Pow(5, float64(i)))
	}
	// fmt.Println(asd)
	return res
}
