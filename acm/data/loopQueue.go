package data

import "fmt"

type loopQueue struct {
	data        []interface{}
	tail, front int
}

func CreateLoopQueue(cap int) *loopQueue {
	return &loopQueue{
		make([]interface{}, cap+1),
		0,
		0,
	}
}

func (q *loopQueue) EnQueue(e interface{}) {
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.GetCap() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
}

func (q *loopQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列为空！", q)
	}
	ret := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	//fmt.Println(q.GetSize(),"`````````",q.GetCap())
	if (q.GetSize() == q.GetCap()/4) && q.GetCap()/2 != 0 {
		*q = q.resize(q.GetCap() / 2)
	}
	return ret
}

func (q *loopQueue) GetFront() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列为空！", q)
	}
	return q.data[q.front]
}

func (q *loopQueue) resize(newCap int) loopQueue {
	//fmt.Println("---------do resize ---------")
	nq := CreateLoopQueue(newCap)
	*nq = nq.copyQueue(q)
	*q = *nq
	return *q
}
func (nq *loopQueue) copyQueue(q *loopQueue) loopQueue {
	for i := 0; i < q.GetSize(); i++ {
		nq.data[i] = q.data[(i+q.front)%len(q.data)]
	}
	nq.front, nq.tail = 0, q.GetSize()
	return *nq
}

func (q *loopQueue) GetCap() int {
	return len(q.data) - 1
}

func (q *loopQueue) GetSize() int {
	return (q.tail - q.front + q.GetCap()) % q.GetCap()
}

func (q *loopQueue) IsEmpty() bool {
	return q.front == q.tail
}

func (q *loopQueue) String() string {
	str := fmt.Sprint("Queue:")
	str += fmt.Sprint(("front ["))
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		if q.data[i]==nil {
			continue
		}
		str += fmt.Sprint((*q).data[i])
		if (i+1)%len(q.data) != q.tail {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] tail")
	return str
}
