package data

import (
	"time"
	"math/rand"
)
type student struct {
	name  string
	score int
}
func TestQueue(q Queue,op int) time.Duration {
	startTime:=time.Now()
	for i:=0;i< op;i++  {
		q.EnQueue(student{"test",rand.Intn(10000)})
	}
	for i:=0;i< op-1;i++  {
		q.DeQueue()
	}

	return time.Since(startTime)
}

