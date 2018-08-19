package lintcode

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	// Write your code here
	exp := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)
	for _, value := range words {
		str := ""
		for _, char := range value {
			if char != 34 {
				str += fmt.Sprint(exp[char-97])
			}
		}
		if str != "" {
			m[str] = 0
		}

	}
	return len(m)
}
