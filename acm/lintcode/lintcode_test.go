package lintcode

import (
	"testing"
)

func TestProblem(t *testing.T) { //lintcode 1013
	//num1 := []int{1, 2, 2, 1}
	//num2 := []int{2, 2}
	//fmt.Println(intersection2(num1, num2))

	topKFrequentWords([]string{
		"yes", "lint", "code",
		"yes", "code", "baby",
		"you", "baby", "chrome",
		"safari", "lint", "code",
		"body", "lint", "code",
	}, 8)
	//coinProblem(100,30)
	//ExamplePriorityQueue()
}
