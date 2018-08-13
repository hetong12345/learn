package main

import (
	"./data"
	"fmt"
	"math/rand"
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
	n:=1000
	for i := 0; i<n; i++ {
		bst.Add(&data.Student{
			Name:"asd",
			Score:rand.Intn(10000),
		})
	}
	al:=data.CreateArray(1000)
	for !bst.IsEmpty() {
		al.AddLast(bst.RemoveMin())
	}
	fmt.Println(al.String())
}

func cal()  {
	fmt.Println(.1 + .2)
	var a float64 = .1
	var b float64 = .2
	fmt.Println(a + b)
	fmt.Printf("%.54f\n", .1 + .2)
}
