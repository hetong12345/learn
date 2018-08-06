package data

import (
	"fmt"
)

type nodeList struct {
	dummyHead node
	size      int
}
type node struct {
	value interface{}
	next  *node
}

//func (n *node)String()string  {
//	//n.value
//	return ""
//}
func CreateList() *nodeList {
	return &nodeList{
		dummyHead: node{
			value: nil,
			next:  nil,
		},
		size: 0,
	}
}
func (l *nodeList) GetSize() int {
	return l.size
}
func (l *nodeList) IsEmpty() bool {
	return l.size == 0
}


func (l *nodeList) AddHead(e interface{}) {
	l.dummyHead.next = &node{
		value: e,
		next:  l.dummyHead.next,
	}

	l.size++
}

func (l *nodeList) AddLast(e interface{}) {
	l.AddIndex(l.size, e)
}
func (l *nodeList) AddIndex(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("fei fa wei zhi")
	}

	prev := &l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next

	}
	newNode := node{
		value: e,
		next:  prev.next,
	}
	fmt.Println(newNode)
	prev.next = &newNode
	fmt.Println("```````````",*l.dummyHead.next)
	l.size++

}
func (l *nodeList) Get(index int) interface{}{
	if index < 0 || index > l.size {
		panic("fei fa wei zhi")
	}
	current := l.dummyHead.next
	for i:=0;i<index ;i++  {
		current=current.next
	}
	return current.value
}
func (l *nodeList) GetFirst() interface{}{
	return l.Get(0)
}
func (l *nodeList) GetLast() interface{}{
	return l.Get(l.size-1)
}
func (l *nodeList) Set(index int,e interface{}) {
	if index < 0 || index > l.size {
		panic("fei fa wei zhi")
	}
	current := l.dummyHead.next
	for i:=0;i<index ;i++  {
		current=current.next
	}
	current.value=e
}
func (l *nodeList) SetFirst(e interface{}) {
	l.Set(0,e)
}
func (l *nodeList) SetLast(e interface{}) {
	l.Set(l.size-1,e)
}
func (l *nodeList)Contains(e interface{}) bool {
	current := l.dummyHead.next
	for current.value!=nil {
		if current.value==e {
			return true
		}
		current=current.next
	}
	return false
}
func (l *nodeList)String() string {
	str := fmt.Sprint("List:")
	str += fmt.Sprint("Head [")
	current := l.dummyHead.next
	for current.value!=nil {
		str+=fmt.Sprint(current.value)
		if current.next!=nil {
			str += fmt.Sprint(",")
			current=current.next
		}
	}
	str += fmt.Sprint("]")
	return str
}
