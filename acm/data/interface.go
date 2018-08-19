package data

type Queue interface {
	//队列
	DeQueue() interface{}
	EnQueue(e interface{})
	GetFront() interface{}
	GetSize() int
	IsEmpty() bool
	String() string
}
type List interface {
	//链表
	RemoveFirst() interface{}
	AddHead(e interface{})
	GetSize() int
	IsEmpty() bool
	Contains(e interface{}) bool
	String() string
}

type Stack interface {
	//栈
	GetSize() int
	IsEmpty() bool
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	String() string
}

type Tree interface {
	//树
	GetSize() int
	IsEmpty() bool
	Add(value Comparable)
	Remove(value Comparable)
	Contains(value Comparable) bool
	Min() Comparable
	Max() Comparable
	String() string
}
type Comparable interface {
	Compare(c2 Comparable) int
}

type Set interface {
	//栈
	GetSize() int
	IsEmpty() bool
	Add(value Comparable)
	Remove(value Comparable)
	Contains(value Comparable) bool
}

type Map interface {
	Add(k interface{}, v interface{})
	Remove(k interface{}) interface{}
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
	Set(key interface{}, newValue interface{})
	GetSize() int
	IsEmpty() bool
}
