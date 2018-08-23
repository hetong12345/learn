package lintcode

import (
	"fmt"
	// "strconv"
)

//noinspection ALL
func Oadd(a int64, b int64) int64 { //不使用加法，计算两数和
	// fmt.Println("test")
	r1 := a ^ b
	r2 := a & b
	// b = r2 << 1
	// fmt.Println(b)
	if r2 != 0 {
		fmt.Println(r1, r2<<1)
		return Oadd(r1, r2<<1)
	}
	// var data1 string = strconv.FormatInt(r1, 2)
	// var data2 string = strconv.FormatInt(r2, 2)
	// fmt.Println(strconv.FormatInt(a, 2), strconv.FormatInt(b, 2), data1, data2)
	// fmt.Println(strconv.FormatInt(77, 2))
	return r1
}
