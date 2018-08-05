package data

type Queue interface {
	DeQueue() interface{}
	EnQueue(e interface{})
	GetFront() interface{}
	GetSize() int
	GetCap() int
	IsEmpty() bool
	String() string
}
