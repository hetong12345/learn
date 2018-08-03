package data

import "fmt"

type queue struct {
	a arr
}

func CreateQueue(cap int) *queue {
	return &queue{*CreateArray(cap)}
}

func (q *queue) GetSize() int {
	return q.a.GetSize()
}

func (q *queue) IsEmpty() bool {
	return q.a.IsEmpty()
}

func (q *queue) EnQueue(e interface{}) {
	q.a.AddLast(e)
}

func (q *queue) DeQueue() interface{} {
	return q.a.RemoveFirst()
}

func (q *queue) GetFront() interface{} {
	return q.a.GetFirst()
}
func (q *queue)String()string  {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint(("front ["))
	for i := 0; i < q.a.size; i++ {
		str += fmt.Sprint((*q).a.data[i])
		if i != q.a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] tail")
	//fmt.Println(str)
	return str
}
