package main

import (
	"fmt"
)

var x = 3

func a() {
	x := 1
	return this.x
}

func b() {
	x := 2
}

func c() {
	a()
}
func main() {
	fmt.Println(x)
	fmt.Println(a())
}

func func_name() {

}
