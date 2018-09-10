package data

type MapTrie struct {
	root *trieNode
	size int
}

type trieNode struct {
	isWord bool
	next   map[rune]*trieNode
}

func CreateMapTrie() *MapTrie {
	return &MapTrie{
		root: &trieNode{
			isWord: false,
			next:   make(map[rune]*trieNode),
		},
		size: 0,
	}
}

func (mt *MapTrie) GetSize() int {
	return mt.size
}

func (mt *MapTrie) Add(word Comparable) { //非递归写法 todo 完成add的递归写法
	cur := mt.root
	char := []rune(string(word.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			//newNode :=createTrieNode()
			//fmt.Println(newNode)
			cur.next[value] = &trieNode{
				isWord: false,
				next:   map[rune]*trieNode{},
			}
		}
		cur = cur.next[value]
	}
	if !cur.isWord {
		cur.isWord = true
		mt.size++
	}
}
func (mt *MapTrie) Contains(word Comparable) bool {
	cur := mt.root
	char := []rune(string(word.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return false
		}
		cur = cur.next[value]
	}
	return cur.isWord
}
func (mt *MapTrie) IsPrefix(prefix Comparable) bool {
	cur := mt.root
	char := []rune(string(prefix.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return false
		}
		cur = cur.next[value]
	}
	return true
}

func (mt *MapTrie) Remove(value Comparable) {
	panic("implement me")
}

func (mt *MapTrie) IsEmpty() bool {
	return mt.size == 0
}

func (mt *MapTrie) String() string {
	panic("implement me")
}
