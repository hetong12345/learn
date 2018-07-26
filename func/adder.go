package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	//panic("implement me")
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

//TODO

func main() {
	f := fibnacci()
	printFile(f)
}
func fibnacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func printFile(reader io.Reader) {
	scan := bufio.NewScanner(reader)
	for scan.Scan() {
		println(scan.Text())
	}
}
