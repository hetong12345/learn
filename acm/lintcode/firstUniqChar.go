package lintcode

func firstUniqChar(s string) int { //leetcode 387
	freq := make([]int, 26)
	for _, value := range s {
		freq[value-'a']++
	}

	for key, value := range s {
		if freq[value-'a'] == 1 {
			return key
		}
	}
	return -1
}
