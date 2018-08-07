package data

type Queue interface {
	DeQueue() interface{}
	EnQueue(e interface{})
	GetFront() interface{}
	GetSize() int
	IsEmpty() bool
	String() string
}
type List interface {
	RemoveFirst() interface{}
	AddHead(e interface{})
	GetSize() int
	IsEmpty() bool
	String() string
}

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	String() string
}
