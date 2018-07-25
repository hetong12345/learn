package data

import (
	"fmt"
)

type Array struct {
	data []int
	size int
}
var arr=Array{}

func main() {
	scores := [...]int{100, 56, 56}
	fmt.Println(scores)

}
//func NewArray(cap int) {
//	Array.Remove()
//	arr.data=[cap]int
//	arr.size =0
//}
//func (*Array)GetSize() int {
//	return arr.size
//}
//func (*Array)GetCap() int {
//	return len(arr.data)
//}
//func (*Array)IsEmpty() bool {
//	return arr.size == 0
//}
//func (*Array)AddLast(e int) {
//	arr.Add(arr.size,e)
//}
//func (*Array)AddFirst(e int)  {
//	arr.Add(0,e)
//}
//func (*Array)Add(index int, e int) {
//	if arr.size == len(arr.data) {
//		panic("out of index")
//	}
//	if index < 0 || index > arr.size {
//		panic("fei fa cha ru")
//	}
//	for i:=arr.size-1;i>=index ; i-- {
//		arr.data[i+1]=arr.data[i]
//	}
//	arr.data[index]=e
//	arr.size++
//}
//func (*Array)ToString()  {
//	fmt.Printf("Array:size=%d,cap=%d\n",arr.size,len(arr.data))
//	fmt.Print("[")
//	for i:=0; i<arr.size; i++ {
//		fmt.Print(arr.data[i])
//		if i!=arr.size-1 {
//			fmt.Print(",")
//		}
//	}
//	fmt.Print("]")
//}
//
//func (*Array)Get(index int ) int {
//	if index < 0 || index > arr.size {
//		panic("fei fa cha ru")
//	}
//	return arr.data[index]
//}
//func (*Array)Set(index int,e int )  {
//	if index < 0 || index > arr.size {
//		panic("fei fa cha ru")
//	}
//	arr.data[index]=e
//}
//func(*Array) Contains(e int ) bool {
//	for i:=0;i<arr.size;i++ {
//		if arr.data[i]==e {
//			return true
//		}
//	}
//	return false
//}
//func (*Array)Find(e int) int {
//	for i:=0;i<arr.size;i++ {
//		if arr.data[i]==e {
//			return i
//		}
//	}
//	return -1
//}
//func (*Array)Remove(index int) int {
//	if index < 0 || index > arr.size {
//		panic("fei fa cha ru")
//	}
//	ret:=arr.data[index]
//	for i:=index+1;i<arr.size ;i++  {
//		arr.data[i-1]=arr.data[i]
//	}
//	arr.size--
//	return ret
//}
//func (*Array)RemoveFirst() int {
//	return arr.Remove(0)
//}
//func(*Array) RemoveLast() int {
//	return arr.Remove(arr.size)
//}
//func(*Array) RemoveElement(e int)  {
//	index := arr.Find(e)
//	if index!=-1 {
//		arr.Remove(index)
//	}
//}