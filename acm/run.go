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
	q:=data.CreateQueue(5)
	for i:=0;i<10 ;i++  {
		q.EnQueue(student{"fasas",i})
	}
	fmt.Println(q.String())

	q.DeQueue()
	fmt.Println(q.String())
}
