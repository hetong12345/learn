package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm/data"
)

func main() {
	//listMap := data.CreateListMap()
	//treeMap := data.CreateTreeMap()
	//listSet := data.CreateListSet()
	//treeSet := data.CreateTreeSet()
	maxHeap := data.CreateMaxHeap()
	fmt.Println(data.TestHeap(maxHeap, 1000000))

	//fmt.Println(data.TestMap(listMap).String())
	//fmt.Println(data.TestMap(treeMap).String())
	//fmt.Println(data.TestSet(listSet).String())
	//fmt.Println(data.TestSet(treeSet).String())

}

func qqq(a ...string) {
	fmt.Println(a)
}
