package acm

import (
	// "fmt"
	"math"
	// "strconv"
)

func Ofactorial(num int64) int64 {
	// fmt.Println(num / 5.0)
	var tmp int64 = num
	var n int
	var res int64
	var i int = 1
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