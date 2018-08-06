package main

import (
	"./data"
	"fmt"
)

type student struct {
	name  string
	score int
}

func main() {
	ls := data.CreateListStack()
	as:=data.CreateArrayStack(10)

	op:=100000000


	fmt.Println(data.TestStack(ls,op))
	fmt.Println(data.TestStack(as,op))
}
