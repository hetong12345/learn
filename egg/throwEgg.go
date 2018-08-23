package main

import (
	"fmt"
)

var floorSum int
var floorbest int
var nextFloor = 1
var operaSum = 0

type egg struct {
	num   int
	statu bool
}

var egg1 = &egg{1, false}
var egg2 = &egg{2, false}

func main() {
	floorSum = 10000
	floorbest = 7491
	fmt.Println("初始化完成，楼层高度为：", floorSum, "最佳楼层为：", floorbest)
	fmt.Println("----------------开始方法一------------------")
	fmt.Println("经计算最佳楼层是", algo1()-1)
	fmt.Println("计算花费步数是：", operaSum)
	fmt.Println("----------------开始方法二------------------")
	fmt.Println("经计算最佳楼层是", algo2()-1)
	fmt.Println("计算花费步数是：", operaSum)
}

func algo1() int { //解决算法，提供下层楼数
	nextFloor = 1
	operaSum = 0
	for {
		egg1.statu = isBroken(nextFloor)
		if egg1.statu {
			break
		}
		nextFloor++
		operaSum++
	}
	return nextFloor
}

func algo2() int {
	nextFloor = 1
	operaSum = 0

	var urange = 500

	for {
		egg1.statu = isBroken(nextFloor)
		if egg1.statu {
			break
		}
		nextFloor += urange
		operaSum++
	}
	nextFloor -= urange
	for {
		egg2.statu = isBroken(nextFloor)
		if egg2.statu {
			break
		}
		nextFloor++
		operaSum++
	}
	operaSum++
	return nextFloor
}

func isBroken(floor int) bool { //判断结果
	if floor <= floorbest {
		return false
	} else {
		return true
	}
}
