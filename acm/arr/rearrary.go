package arr

import (
	"fmt"
	"log"
)

type arr struct {
	data []int
	size int
}

func Create(cap int) arr {
	a := []int{}
	for i := 0; i < cap; i++ {
		a = append(a, 0)
	}
	return arr{
		data: a,
		size: 0,
	}
}
func (a *arr) GetSize() int {
	return a.size
}
func (a *arr) GetCap() int {
	return len(a.data)
}
func (a *arr) AddLast(e int) bool {
	if a.Add(a.size, e) {
		fmt.Println(a)
		return true
	}
	return false
}
func (a *arr) AddFirst(e int) bool {
	if a.Add(0, e) {
		return true
	}
	return false
}
func (a *arr) Add(index int, e int) bool {
	if a.size == len(a.data) {
		log.Println("mei you kong jian ")
		return false
	}
	if index < 0 || index > a.size {
		log.Println("fei fa weizhi ")
		return false
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
		fmt.Println(a.data[i+1])
	}
	a.data[index] = e
	return true
}
func (a *arr) IsEmpty() bool {
	return a.size == 0
}

//func (arr *arr)IsEmpty() bool {
//	return arr.size==0
//}
//func (arr *arr)IsEmpty() bool {
//	return arr.size==0
//}
