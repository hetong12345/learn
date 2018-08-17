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

	a := data.
		fmt.Println(a.Compare())
	//for _, value := range split {
	//	treeSet.Add(strings.ToLower(value))
	//}
	fmt.Println(len(split))
}
