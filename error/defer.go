package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()
	f := fibnacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(write, f())
	}
}
func fibnacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

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
func main() {
	writeFile("fib.txt")
}
