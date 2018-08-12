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
	Add(Comparable)
	Contains(Comparable) bool
}
type Comparable interface {
	CompareTo(var2 Comparable) int
}
