package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm/data"
)

func main() {
	//listMap := data.CreateListMap()
	treeMap := data.CreateTreeMap()
	//listSet := data.CreateListSet()
	//treeSet := data.CreateTreeSet()

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	lintcode.i

	//fmt.Println(data.TestMap(listMap).String())
	fmt.Println(data.TestMap(treeMap).String())
	//fmt.Println(data.TestSet(listSet).String())
	//fmt.Println(data.TestSet(treeSet).String())

}
