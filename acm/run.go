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
	treeMap.Add(data.Integer(6), 1)
	treeMap.Add(data.Integer(7), 1)
	treeMap.Add(data.Integer(4), 1)
	treeMap.Add(data.Integer(10), 1)
	treeMap.Add(data.Integer(2), 1)
	treeMap.Add(data.Integer(5), 1)
	fmt.Println(treeMap)
	fmt.Println(treeMap.GetSize())
	fmt.Println(treeMap.IsEmpty())

	m := map[int]int{1: 1}

	//fmt.Println(data.TestMap(listMap).String())
	////fmt.Println(data.TestSet(listSet).String())
	//fmt.Println(data.TestSet(treeSet).String())

}
