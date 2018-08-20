package data

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func TestQueue(q Queue, op int) time.Duration {
	startTime := time.Now()
	for i := 0; i < op; i++ {
		q.EnQueue(Student{"test", rand.Intn(10000)})
	}
	for i := 0; i < op-1; i++ {
		q.DeQueue()
	}

	return time.Since(startTime)
}
func TestStack(s Stack, op int) time.Duration {
	startTime := time.Now()
	for i := 0; i < op; i++ {
		s.Push(Student{"test", rand.Intn(10000)})
	}
	for i := 0; i < op-1; i++ {
		s.Pop()
	}

	return time.Since(startTime)
}

func TestSet(s Set) time.Duration {
	startTime := time.Now()

	file := "./acm/pride-and-prejudice.txt"
	//file := "./acm/测试文本"

	word := divideWord(file)
	fmt.Println(len(word))

	for _, value := range word {

		ls := strings.ToLower(value)

		s.Add(Stringer(ls))
	}
	fmt.Println(s.GetSize())

	return time.Since(startTime)
}

func TestMap(m Map) time.Duration {
	startTime := time.Now()

	file := "./acm/pride-and-prejudice.txt"
	//file := "./acm/测试文本"

	word := divideWord(file)
	fmt.Println(len(word))

	for _, value := range word {
		ls := strings.ToLower(value)
		num := m.Get(ls)
		if num == nil {
			m.Add(ls, 1)
		} else {
			m.Set(ls, num.(int)+1)
		}

	}
	fmt.Println(m.GetSize())

	return time.Since(startTime)
}
func divideWord(path string) []string {
	var wordMap []string
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf("读文件错误", err)
	}
	text := string(file)
	//按字母读取，除26个字母（大小写）之外的所有字符均认为是分隔符
	var size = 0
	letterStart := false
	var str = ""
	for _, value := range text {
		if (value >= 65 && value <= 90) || (value >= 97 && value <= 122) {
			letterStart = true
			str += fmt.Sprintf("%c", value)
		} else {
			if letterStart {
				size++
				letterStart = false
				wordMap = append(wordMap, str)
				//fmt.Println(str,wordMap)
				str = ""
			}
		}
	}
	return wordMap
}
