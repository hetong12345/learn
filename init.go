package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm"
)

func main() {
	// var res int64 = acm.Oadd(123312312, 33242343)
	// fmt.Println(res)

	// var fa int64 = acm.Ofactorial(125.0)
	// fmt.Println(fa)

	// var count int = acm.OdigitCounts(1, 1)
	// fmt.Println(count)
	// if len(nums) == 0 {
	// 	return 0
	// } else {

	// 	for i := 0; i < len(nums); i++ {
	// 		fmt.Println(i)
	// 	}
	// 	return 1
	// }
	nums := []int{1, 2, 2}
	a := acm.RemoveDuplicates(nums)
	fmt.Println(a)
}
