package main

import (
	"./data"
	"fmt"
)


func main() {
	aq:=data.CreateArrayQueue(10)
	lq:=data.CreateLoopQueue(10)
	fmt.Println(data.TestQueue(aq,100000))
	fmt.Println(data.TestQueue(lq,100000))
}
