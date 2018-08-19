package main

import (
	"github.com/hetong12345/learn/acm/data"
)

func main() {
	listMap := data.CreateListMap()
	//listSet := data.CreateListSet()

	println(data.TestMap(listMap).String())
	//println(data.TestSet(listSet).String())

}
