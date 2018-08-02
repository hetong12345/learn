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

	a:=data.Create(5)
	for i := 0; i<100;i++  {
		a.AddLast(student{"test",i})
	}
	fmt.Println(a.String())
	for i:=0;i<60 ;i++  {
		a.RemoveLast()
	}
	fmt.Println(a)

}
