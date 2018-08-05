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
	q := data.CreateLoopQueue(7)
	for i:=0;i<15 ;i++  {
		q.EnQueue(student{"name",i})
		fmt.Println(q.String())

		if i>5 {
			q.DeQueue()
			fmt.Println(q.String())
		}
	}
}
