package lintcode

type MapSum struct {
	root *trieNode
	size int
}

type trieNode struct {
	isWord bool
	next   map[rune]*trieNode
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		root: &trieNode{
			isWord: false,
			next:   make(map[rune]*trieNode),
		},
		size: 0,
	}
}

func (this *MapSum) Insert(key string, val int) {

}

func (this *MapSum) Sum(prefix string) int {

}
