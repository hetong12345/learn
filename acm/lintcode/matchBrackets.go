package lintcode

import (
	"../data"
)

func IsValidParentheses(s string) bool {//用栈实现括号匹配
	stack := data.CreateArrayStack(5)
	for _, v := range s {
		if string(v) == "{" || string(v) == "[" || string(v) == "(" {
			stack.Push(string(v))
		} else {
			if stack.IsEmpty() {
				return false
			}
			switch stack.Pop() {
			case "{":
				if string(v) != "}" {
					return false
				}
			case "[":
				if string(v) != "]" {
					return false
				}
			case "(":
				if string(v) != ")" {
					return false
				}
			}
		}
	}
	return stack.IsEmpty()
}
