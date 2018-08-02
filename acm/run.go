package main

import (
	"./lintcode"
)

type student struct {
	name  string
	score int
}

func main() {
	println(lintcode.KthLargest(4 ,[]int{5,1,6,8,3,1,78,8}))

}
