package main

import (
	"./lintcode"
	"fmt"
	"strings"
	"math"
)

type student struct {
	name  string
	score int
}

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	example := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	lintcode.ToeplitzMatrix(example)

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex 实现了 Abser
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v
	fmt.Println(a.Abs())
	ad:=IPAddr{127, 0, 0, 1}
	//fmt.Println(ad[0]+"."+ad[1]+"."+ad[2]+".")
	fmt.Print(ad.String())

}

type IPAddr [4]byte
func (ad *IPAddr)String() string {
	fmt.Println(string(ad[0]))
	a:=string(ad[0])+"."+string(ad[1])+"."+string(ad[2])+"."+string(ad[3])
	return a
}
func Pic(dx, dy int) [][]uint8 {
	ret := [][]uint8{}
	for i := 0; i < dy; i++ {
		x := make([]uint8, dx)
		for j := 0;j<dx ;j++ {
			x[j]=uint8(i%j+1)
		}
		ret = append(ret, x)
	}
	return ret
}
func Sqrt(x float64) float64 {
	var z = 1.0
	i := 0
	for {
		newz := z - (z*z-x)/(2*z)
		i++
		fmt.Println(z)
		if newz-z < 0.0000000000001 && newz-z > -0.0000000000001 {
			break
		}
		z = newz
	}
	fmt.Println(i)
	return z
}
func WordCount(s string) map[string]int {
	a:=strings.Fields(s)
	ret:=make(map[string]int)
	for i := 0; i < len(a); i++ {
		if v,ok:=ret[a[i]];ok {
			(ret)[a[i]]=v+1
		}else {
			(ret)[a[i]]=1
		}
	}
	return ret
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func fibonacci() func() int {
	a,b :=0,1
	return func() int {
		a , b= b, a+b
		return b
	}
}

type Abser interface {
	Abs() float64
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}