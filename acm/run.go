package main

import (
	"errors"
	"fmt"
	"github.com/hetong12345/learn/acm/algorithm"
	"net/http"
	"runtime"
	"sort"
	"strconv"
	"sync"
)

func main() {
	//n := 100000
	fmt.Println(algorithm.Comp("1.2.5c", "1.2.6b"))
	//arr1 := algorithm.CreateRandomArray(n, 0, n)
	//arr2 := algorithm.CreateRandomArray(n, 0, n)
	//arr3 := algorithm.CreateRandomArray(n, 0, n)
	//arr4 := algorithm.CreateRandomArray(n, 0, n)

	//fmt.Println(algorithm.TestSort(algorithm.SelectionSort, arr1))
	//fmt.Println(algorithm.TestSort(algorithm.InsertionSort, arr2))
	//fmt.Println(algorithm.TestSort(algorithm.InsertionSort, arr3))
	//fmt.Println(algorithm.TestSort(algorithm.BubbleSort, arr4))
}
func comp() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	fmt.Println(sn1)
	fmt.Println(sn2)
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	} else {
		fmt.Println("error")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}

}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println(x)
	fmt.Println("non-empty interface")
}
func testCond() {
	//locker
	//cd:=sync.NewCond()
}
func testGroup() {
	wg := sync.WaitGroup{}
	//var wg sync.WaitGroup
	var urls = []string{
		"https://www.baidu.com/",
		"https://www.aliyun.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func waitGroup() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(4)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func sortSlice() {
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

func convString() {
	//fmt.Println(data.TestMap(listMap).String())
	//fmt.Println(data.TestMap(treeMap).String())
	//fmt.Println(data.TestSet(listSet).String())
	//fmt.Println(data.TestSet(treeSet).String())
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
}

type student struct {
	Name string
	Age  int
}

func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}
	//
	//for k,v:=range m{
	//	println(k,"=>",v.Name)
	//}

	for _, stu := range stus {
		fmt.Println(stu.Name, stu)
		m[stu.Name] = &stu
		fmt.Println(m[stu.Name])
	}

	for k, v := range m {

		fmt.Println(k, "=>", (*v).Name)
	}

}

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
