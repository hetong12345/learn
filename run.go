package main

import (
	"fmt"
	"github.com/hetong12345/learn/single"
)

func main() {
	fmt.Printf("even: %v\n", single.Even(100))
	a, b := single.Swap(100, 200)
	fmt.Printf("swap: %d\t%d\n", a, b)
	fmt.Printf("shifting: %d\n", single.Shifting(100))
	fmt.Printf("nagation: %d\n", single.Nagation(100))
	fmt.Println(1 & 2)
	fmt.Println(2 & 2)
	fmt.Println(3 & 2)
	fmt.Println(4 & 2)
	b = 10 << 1
	fmt.Printf("%v", b)
}
