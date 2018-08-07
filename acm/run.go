package main

import (
	"./data"
	"./lintcode"
	"fmt"
)

type student struct {
	name  string
	score int
}


func main() {
	//ls := data.CreateListStack()
	//as:=data.CreateArrayStack(10)
	//
	op := 100000

	listq := data.CreateListQueue()
	loopq := data.CreateLoopQueue(10)
	aq := data.CreateArrayQueue(10)

	fmt.Println(data.TestQueue(listq, op))
	fmt.Println(data.TestQueue(loopq, op))
	fmt.Println(data.TestQueue(aq, op))

	head:=lintcode.ListNode{
		1,
		&lintcode.ListNode{
			2,
			&lintcode.ListNode{},
		},
	}
	lintcode.RemoveElements(&head,4)
}

