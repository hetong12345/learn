package lintcode

type WordDictionary struct {
	//leetcode 211
	root *trieNode
	size int
}

//type trieNode struct {
//	isWord bool
//	next   map[rune]*trieNode
//}

/** Initialize your data structure here. */
//func Constructor() WordDictionary {
//	return WordDictionary{
//		root: &trieNode{
//			isWord: false,
//			next:   make(map[rune]*trieNode),
//		},
//		size: 0,
//	}
//}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	char := []rune(word)

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			cur.next[value] = &trieNode{
				isWord: false,
				next:   map[rune]*trieNode{},
			}
		}
		cur = cur.next[value]
	}
	if !cur.isWord {
		cur.isWord = true
		this.size++
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search(word, this.root, 0)
}

func (this *WordDictionary) search(word string, node *trieNode, index int) bool {
	if index == len(word) {
		return node.isWord
	}

	char := []rune(word)

	if string(char[index]) != "." {
		if _, ok := node.next[char[index]]; !ok {
			return false
		} else {
			return this.search(word, node.next[char[index]], index+1)
		}
	} else {
		for _, value := range node.next {
			if this.search(word, value, index+1) {
				return true
			}
		}
		return false
	}
}

//return cur.isWord

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
