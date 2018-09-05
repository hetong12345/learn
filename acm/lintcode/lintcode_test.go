package lintcode

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) { //lintcode 1013
	//num1 := []int{1, 2, 2, 1}
	//num2 := []int{2, 2}
	//fmt.Println(intersection2(num1, num2))

	//a := topKFrequentWords([]string{
	//	"yes", "lint", "code",
	//	"yes", "code", "baby",
	//	"you", "baby", "chrome",
	//	"safari", "lint", "code",
	//	"body", "lint", "code",
	//}, 4)
	//fmt.Println(a)
	//fmt.Println(topk([]int{1, 2, 3, 4, 5, 6, 8, 9, 10, 7}, 10))
	//coinProblem(100,30)
	//ExamplePriorityQueue()

	nums := []int{}
	obj := Constructor(nums)
	fmt.Println(obj)
	//param_1 := obj.SumRange(0,2)
	param_2 := obj.SumRange(2, 5)
	//param_3 := obj.SumRange(0,5)
	//fmt.Println(param_1)
	fmt.Println(param_2)
	//fmt.Println(param_3)
}
