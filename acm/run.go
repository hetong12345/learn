package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm/data"
	"sort"
	"strconv"
	"sync"
)

func main() {
	//listMap := data.CreateListMap()
	//treeMap := data.CreateTreeMap()
	//listSet := data.CreateListSet()
	//treeSet := data.CreateTreeSet()

	maxHeap := data.CreateMaxHeap()
	fmt.Println(data.TestHeap(maxHeap, 1000000))

	//fmt.Println(data.TestMap(listMap).String())
	//fmt.Println(data.TestMap(treeMap).String())
	//fmt.Println(data.TestSet(listSet).String())
	//fmt.Println(data.TestSet(treeSet).String())
	s := "asddsa"
	qqq(s)

	//第一种方式 fmt.Sprintf
	//str = fmt.Sprintf("%d", num1)
	//fmt.Printf("str type %T str=%v\n", str, str)
	//
	//str = fmt.Sprintf("%f", num2)
	//fmt.Printf("str type %T str=%v\n", str, str)
	//
	//str = fmt.Sprintf("%t", b)
	//fmt.Printf("str type %T str=%v\n", str, str)
	//
	//str = fmt.Sprintf("%c", ch)
	//fmt.Printf("str type %T str=%v\n", str, str)
	//
	////第二种方式 strconv
	//str = strconv.FormatInt(int64(num1), 10)
	//fmt.Printf("str type %T str=%v\n", str, str)

	var str1 string = "true"
	var b bool

	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "123456"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)

	var str3 string = "123dd456"
	var n2 int64
	n2, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Printf("n1 type %T n1=%v\n", n2, n2)

	var str4 string = "123"
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	n5 := int(n3)
	fmt.Printf("n3 type %T n1=%v\n", n3, n3)
	fmt.Printf("n5 type %T n1=%v\n", n5, n5)

	var str5 string = "123.456"
	var n4 float64
	n4, _ = strconv.ParseFloat(str5, 64)
	fmt.Printf("n1 type %T n1=%v\n", n4, n4)

	type safeMap struct {
		sync.Map
	}
	var sm safeMap
	sm.Store(5, 5)
	fmt.Println(sm)
	value, ok := sm.Load(5)
	fmt.Println(value, ok)

	type Peak struct {
		Name      string
		Elevation int // in feet
	}

	peaks := []Peak{
		{"Aconcagua", 22838},
		{"Denali", 20322},
		{"Kilimanjaro", 19341},
		{"Mount Elbrus", 18510},
		{"Mount Everest", 29029},
		{"Mount Kosciuszko", 7310},
		{"Mount Vinson", 16050},
		{"Puncak Jaya", 16024},
	}

	// does an in-place sort on the peaks slice, with tallest peak first
	sort.Slice(peaks, func(i, j int) bool {
		return peaks[i].Elevation >= peaks[j].Elevation
	})

	// peaks is now sorted
}

func qqq(a interface{}) {
	c, ok := a.(int)
	if ok {
		fmt.Println(c)
	}
	fmt.Println(a)
}
