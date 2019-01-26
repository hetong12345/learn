package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm/algorithm"
)

func main() {

	n := 100000

	//set := data.CreateAvlTreeSet()
	//fmt.Println(data.TestSet(set))

	arr1 := algorithm.CreateRandomArray(n, 0, n)
	arr2 := algorithm.CreateRandomArray(n, 0, n)
	arr3 := algorithm.CreateRandomArray(n, 0, n)
	arr4 := algorithm.CreateRandomArray(n, 0, n)

	fmt.Println(algorithm.TestSort(algorithm.SelectionSort, arr1))
	fmt.Println(algorithm.TestSort(algorithm.InsertionSort, arr2))
	fmt.Println(algorithm.TestSort(algorithm.ShellSort, arr3))
	fmt.Println(algorithm.TestSort(algorithm.BubbleSort, arr4))
}
