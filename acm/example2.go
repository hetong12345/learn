package main

import (
	"fmt"
)

type option1 struct {
	a int64
	b string
	c map[int]string
}

func (o *option1) writeA(a int64) *option1 {
	o.a = a
	return o
}
func (o *option1) writeB(b string) *option1 {
	o.b = b
	return o
}

func (o *option1) writeC(c map[int]string) *option1 {
	o.c = c
	return o
}

func main() {
	op := new(option1)
	op.writeA(int64(1)).writeB("test").writeC(make(map[int]string, 0))

	fmt.Println(op.a, op.b, op.c)
}