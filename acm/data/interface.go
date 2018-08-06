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
type List interface {
	//DeQueue() interface{}

	GetSize() int
	IsEmpty() bool
	String() string
}