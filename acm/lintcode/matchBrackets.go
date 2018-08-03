package lintcode

import (
	"../data"
)
func IsValidParentheses (s string) bool {
	stack:=data.CreateQueue(5)
	for _,v := range s {
		if string(v)=="{"|| string(v)=="["||string(v)=="("{
			stack.Push(string(v))
		}else {
			if stack.IsEmpty() {
				return false
			}
			switch stack.Pop() {
			case "{":
				if string(v)!="}" {
					return false
				}
			case "[":
				if string(v)!="]" {
					return false
				}
			case "(":
				if string(v)!=")" {
					return false
				}
			}
		}
	}
	return stack.IsEmpty()
}