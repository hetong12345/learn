package data

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func TestQueue(q Queue, op int) time.Duration {
	startTime := time.Now()
	for i := 0; i < op; i++ {
		q.EnQueue(Student{"test", rand.Intn(10000)})
	}
	for i := 0; i < op-1; i++ {
		q.DeQueue()
	}

	return time.Since(startTime)
}
func TestStack(s Stack, op int) time.Duration {
	startTime := time.Now()
	for i := 0; i < op; i++ {
		s.Push(Student{"test", rand.Intn(10000)})
	}
	for i := 0; i < op-1; i++ {
		s.Pop()
	}

	return time.Since(startTime)
}

func TestSet(s Set) time.Duration {
	startTime := time.Now()

	file, err := ioutil.ReadFile("./acm/pride-and-prejudice.txt")
	if err != nil {
		fmt.Errorf("读文件错误", err)
	}
	str := string(file)
	split := strings.Split(str, " ")
	fmt.Println(len(split))

	for _, value := range split {
		ls := strings.ToLower(value)

		s.Add(Stringer(ls))
	}
	fmt.Println(s.GetSize())

	return time.Since(startTime)
}
func TestMap(m Map) time.Duration {
	startTime := time.Now()

	file, err := ioutil.ReadFile("./acm/pride-and-prejudice.txt")
	if err != nil {
		fmt.Errorf("读文件错误", err)
	}
	str := string(file)
	split := strings.Split(str, " ")
	fmt.Println(len(split))

	for _, value := range split {
		ls := strings.ToLower(value)
		num := m.Get(ls)
		if num == nil {
			m.Add(ls, 1)
		} else {
			m.Set(ls, num.(int)+1)
		}

	}
	fmt.Println(m.GetSize())

	return time.Since(startTime)
}
