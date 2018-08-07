package data

import (
	"fmt"
)

type NodeList struct {
	dummyHead node
	size      int
}
type node struct {
	value interface{}
	next  *node
}

func CreateNodeList() *NodeList {
	return &NodeList{
		dummyHead: node{
			value: nil,
			next:  nil,
		},
		size: 0,
	}
}

func (nl *NodeList) GetSize() int {
	return nl.size
}

func (nl *NodeList) IsEmpty() bool {
	return nl.size == 0
}

func (nl *NodeList) AddHead(e interface{}) {
	nl.dummyHead.next = &node{
		value: e,
		next:  nl.dummyHead.next,
	}
	nl.size++
}

func (nl *NodeList) AddLast(e interface{}) {
	nl.AddIndex(nl.size, e)
}
func (nl *NodeList) AddIndex(index int, e interface{}) {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	prev := &nl.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	newNode := node{
		value: e,
		next:  prev.next,
	}
	prev.next = &newNode
	nl.size++
}
func (nl *NodeList) Get(index int) interface{} {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	current := nl.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}
func (nl *NodeList) GetFirst() interface{} {
	return nl.Get(0)
}
func (nl *NodeList) GetLast() interface{} {
	return nl.Get(nl.size - 1)
}
func (nl *NodeList) Set(index int, e interface{}) {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	current := nl.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = e
}
func (nl *NodeList) SetFirst(e interface{}) {
	nl.Set(0, e)
}
func (nl *NodeList) SetLast(e interface{}) {
	nl.Set(nl.size-1, e)
}
func (nl *NodeList) Contains(e interface{}) bool {
	current := nl.dummyHead.next
	for current.value != nil {
		if current.value == e {
			return true
		}
		current = current.next
	}
	return false
}
func (nl *NodeList) Remove(index int) interface{} {
	if index < 0 || index > nl.size {
		panic("fei fa wei zhi")
	}
	prev := &nl.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	ret := prev.next
	prev.next = ret.next
	ret.next = nil

	nl.size--
	return ret.value
}
func (nl *NodeList) RemoveFirst() interface{} {
	return nl.Remove(0)
}
func (nl *NodeList) RemoveLast() interface{} {
	return nl.Remove(nl.size - 1)
}
func (nl *NodeList) String() string {
	str := fmt.Sprint("List:")
	str += fmt.Sprint("Head ")
	current := nl.dummyHead.next
	for {
		str += fmt.Sprint(current.value)
		if current.next != nil {
			str += fmt.Sprint("->")
			current = current.next
		} else {
			break
		}
	}
	str += fmt.Sprint("->nil")
	return str
}
