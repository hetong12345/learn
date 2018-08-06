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
	nl:=data.CreateList()
	for i:=0;i<3 ;i++  {
		nl.AddHead(student{
			name:"asd",
			score:i,
		})
	}
	//for i:=0;i<20 ;i++  {
	//	fmt.Println(nl)
	//}
	fmt.Println(nl.String())
}
