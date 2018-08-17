package main

import (
	"fmt"
	"github.com/hetong12345/learn/acm/data"
	"io/ioutil"
	"strings"
)

func main() {
	//file, err := ioutil.ReadFile("./acm/a-tale-of-two-cities.txt")
	file, err := ioutil.ReadFile("./acm/pride-and-prejudice.txt")
	if err != nil {
		fmt.Errorf("读文件错误", err)
	}
	str := string(file)
	split := strings.Split(str, " ")
	//treeSet := data.CreateTreeSet()
	fmt.Println(data.Ints(5, 6))
	a := data.TwoStudent{
		data.Student{
			Name:  "zhang san",
			Score: 55,
		},
		data.Student{
			Name:  "renketong",
			Score: 99,
		},
	}
	fmt.Println(a.Compare())
	//for _, value := range split {
	//	treeSet.Add(strings.ToLower(value))
	//}
	fmt.Println(len(split))
}
