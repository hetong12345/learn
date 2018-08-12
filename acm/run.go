package main

import (
	"./data"
	"fmt"
)

type student struct {
	name  string
	score int
}

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	bst := data.CreateBinarySearchTree()
	nums := []int{5, 365, 13, 2, 65, 73, 66, 22, 77, 99}
	for _, value := range nums {
		bst.Add(&data.Student{
			Name:  "asd",
			Score: value,
		})
	}

	fmt.Println(bst.Min())
	fmt.Println(bst.Max())
	//fmt.Println(bst.String())
}
