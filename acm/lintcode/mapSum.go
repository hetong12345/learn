package lintcode

type MapSum struct {
	root *trieNode
}

type trieNode struct {
	value int
	next  map[rune]*trieNode
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{
		root: &trieNode{
			value: 0,
			next:  make(map[rune]*trieNode),
		},
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	char := []rune(key)

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			cur.next[value] = &trieNode{
				value: 0,
				next:  map[rune]*trieNode{},
			}
		}
		cur = cur.next[value]
	}
	cur.value = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	char := []rune(prefix)

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return 0
		}
		cur = cur.next[value]
	}

	return this.sum(cur)
}

func (this *MapSum) sum(node *trieNode) int {

	ret := node.value

	for _, value := range node.next {
		ret += this.sum(value)
	}
	return ret
}
