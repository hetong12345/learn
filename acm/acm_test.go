package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm/data"
	"math/rand"
	"testing"
	"time"
)

func TestBst(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	bst := data.CreateBinarySearchTree()
	for i := 0; i < 1000; i++ {
		bst.Add(&data.Student{
			Name:  "asd",
			Score: rand.Intn(10000),
		})
	}
	il := data.CreateArray(1000)
	for !bst.IsEmpty() {
		il.AddLast(bst.RemoveMax())
	}
	for i := 1; i < il.GetSize(); i++ {
		v1 := il.Get(i - 1).(*data.Student)
		v2 := il.Get(i).(*data.Student)
		if v1.Compare(v2) < 0 {
			t.Error("大小错位")
		}
	}
	t.Log("ok!")
}

func TestRemove(t *testing.T) {
	bst := data.CreateBinarySearchTree()

	bst.Add(&data.Student{
		Name:  "asd",
		Score: 4,
	})
	bst.Add(&data.Student{
		Name:  "asd",
		Score: 1,
	})
	bst.Add(&data.Student{
		Name:  "asd",
		Score: 3,
	})
	bst.Add(&data.Student{
		Name:  "asd",
		Score: 8,
	})
	bst.Add(&data.Student{
		Name:  "asd",
		Score: 6,
	})
	fmt.Println(bst)
	bst.Remove(&data.Student{
		Name:  "asd",
		Score: 4,
	})
	fmt.Println(bst)
}

func TestBstInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	bst := data.CreateBinarySearchTree()
	for i := 0; i < 1000; i++ {
		bst.Add(data.Integer(rand.Intn(10000)))
	}
	il := data.CreateArray(1000)
	for !bst.IsEmpty() {
		il.AddLast(bst.RemoveMax())
	}
	fmt.Println(il)
	for i := 1; i < il.GetSize(); i++ {
		v1 := il.Get(i - 1).(data.Integer)
		v2 := il.Get(i).(data.Integer)
		if v1.Compare(v2) < 0 {
			t.Error("大小错位")
		}
	}
	t.Log("ok!")
}
