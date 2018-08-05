package data

import "fmt"

type arrayQueue struct {
	a arr
}

func CreateArrayQueue(cap int) *arrayQueue {
	return &arrayQueue{*CreateArray(cap)}
}

func (q *arrayQueue) GetSize() int {
	return q.a.GetSize()
}

func (q *arrayQueue) IsEmpty() bool {
	return q.a.IsEmpty()
}

func (q *arrayQueue) EnQueue(e interface{}) {
	q.a.AddLast(e)
}

func (q *arrayQueue) DeQueue() interface{} {
	return q.a.RemoveFirst()
}

func (q *arrayQueue) GetFront() interface{} {
	return q.a.GetFirst()
}
func (q *arrayQueue) GetCap() int {
	return q.a.GetCap()
}
func (q *arrayQueue) String() string {
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
