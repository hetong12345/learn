package lintcode

import "fmt"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	ret := ""
	var tmp byte
	for i := 0; i < len(strs[0]); i++ {
		for _, value := range strs {
			if value == "" {
				return ""
			}
			if value == strs[0] {
				tmp = strs[0][i]
			}
			if i >= len(value) {
				return ret
			}
			if []byte(value)[i] != tmp {
				fmt.Println(value, i)
				return ret
			}
			tmp = []byte(value)[i]
			fmt.Println(string(tmp))
		}
		ret += string(tmp)
	}
	fmt.Println(strs)
	return strs[0]
}
