package algorithm

import (
	"strconv"
	"strings"
)

func Comp(v1, v2 string) int { //比较v1 v2 的版本号返回 v1大返回1 v2大返回-1 相同返回0
	if v1 == "" || v2 == "" {
		panic("传入有误")
	}
	if v1 == v2 {
		return 0
	}
	a1 := strings.FieldsFunc(v1, spiltNet) //以 “.”分割传入的版本号
	a2 := strings.FieldsFunc(v2, spiltNet)
	for key, value := range a1 {
		if key > len(a2) { //如果前面位数都相同，a2没有后续，则a1大
			return 1
		}
		if value == a2[key] { //如果当前位数相同，则进入下层
			continue
		}
		if key == len(a1)-1 && key == len(a2)-1 {
			return compStr(value, a2[key])
		}
		n1, _ := strconv.Atoi(value)
		n2, _ := strconv.Atoi(a2[key])
		if n1 > n2 {
			return 1
		} else if n2 > n1 {
			return -1
		}
	}
	return -1
}
func compStr(s1 string, s2 string) int {
	for key, value := range s1 {
		if value != int32(s2[key]) {
			if value < int32(s2[key]) {
				return -1
			} else {
				return 1
			}
		}
	}
	return 0
}

func spiltNet(char rune) bool {
	if char == '.' {
		return true
	}
	return false
}
