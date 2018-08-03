package data

import "fmt"

type stack struct {
	a arr
}

func CreateStack(cap int) *stack {
	return &stack{*CreateArray(cap)}
}

func (s *stack) GetSize() int {
	return s.a.GetSize()
}

func (s *stack) IsEmpty() bool {
	return s.a.IsEmpty()
}

func (s *stack) Push(e interface{}) {
	s.a.AddLast(e)
}

func (s *stack) Pop() interface{} {
	return s.a.RemoveLast()
}

func (s *stack) Peek() interface{} {
	return s.a.GetLast()
}
func (s *stack)String()string  {
	str := fmt.Sprint("Stack:")
	str += fmt.Sprint(("["))
	for i := 0; i < s.a.size; i++ {
		str += fmt.Sprint((*s).a.data[i])
		if i != s.a.size-1 {
			str += fmt.Sprint(",")
		}
	}
	str += fmt.Sprint("] top")
	//fmt.Println(str)
	return str
}
