package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	println(t.Value)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	return true
}

func main() {
	ch:=make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println("123")
	select {

	}
}
