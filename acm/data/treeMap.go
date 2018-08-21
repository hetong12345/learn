package data

import "fmt"

type TreeMap struct {
	root *treeMapNode
	size int
}

type treeMapNode struct {
	key         Comparable
	value       interface{}
	left, right *treeMapNode
}

func CreateTreeMap() *TreeMap {
	return &TreeMap{
		root: nil,
		size: 0,
	}
}

func (tm *TreeMap) Add(k interface{}, v interface{}) {
	tm.root = tm.add(tm.root, k, v)
}
func (tm *TreeMap) add(node *treeMapNode, k interface{}, v interface{}) *treeMapNode {
	if node == nil {
		tm.size++
		return &treeMapNode{key: k.(Comparable), value: v, left: nil, right: nil}
	}
	if k.(Comparable).Compare(node.key) < 0 {
		node.left = tm.add(node.left, k, v)
	} else if k.(Comparable).Compare(node.key) > 0 {
		node.right = tm.add(node.right, k, v)
	} else {
		node.value = v
	}
	return node
}
func (tm *TreeMap) Remove(k interface{}) interface{} {
	ret := tm.Get(k)
	tm.root = tm.remove(tm.root, k)
	return ret
}
func (tm *TreeMap) remove(node *treeMapNode, k interface{}) *treeMapNode {
	if node == nil {
		return nil
	}
	if k.(Comparable).Compare(node.key) < 0 {
		node.left = tm.remove(node.left, k)
		return node
	} else if k.(Comparable).Compare(node.key) > 0 {
		node.right = tm.remove(node.right, k)
		return node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			tm.size--
			return right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			tm.size--
			return left
		} //todo complete this part

		ret := tm.min(node.right)
		ret.right = tm.remove(node.right, ret.key)
		ret.left = node.left
		node.left, node.right = nil, nil
		return ret
	}
	//if k.Compare(node.value) < 0 {
	//	node.left = bst.remove(node.left, value)
	//	return node
	//} else if value.Compare(node.value) > 0 {
	//	node.right = bst.remove(node.right, value)
	//	return node
	//} else {
	//	if node.left == nil {
	//		right := node.right
	//		node.right = nil
	//		bst.size--
	//		return right
	//	} else if node.right == nil {
	//		left := node.left
	//		node.left = nil
	//		bst.size--
	//		return left
	//	}
	//
	//	ret := min(node.right)
	//	ret.right = removeMin(node.right)
	//	ret.left = node.left
	//	node.left, node.right = nil, nil
	//	return ret
	//
	//}
}

func (tm *TreeMap) min(node *treeMapNode) *treeMapNode {
	if node.left == nil {
		return node
	}
	return tm.min(node.left)
}
func (tm *TreeMap) max(node *treeMapNode) *treeMapNode {
	if node.right == nil {
		return node
	}
	return tm.max(node.right)
}
func (tm *TreeMap) Contains(key interface{}) bool {
	return tm.contains(tm.root, key)
}
func (tm *TreeMap) contains(node *treeMapNode, key interface{}) bool {
	if node == nil {
		return false
	} else if node.key == key {
		return true
	}
	if node.key.Compare(key.(Comparable)) < 0 {
		return tm.contains(node.left, key)
	} else {
		return tm.contains(node.right, key)
	}
}
func (tm *TreeMap) Get(key interface{}) interface{} {
	return tm.get(tm.root, key)
}
func (tm *TreeMap) get(node *treeMapNode, key interface{}) interface{} {
	if tm.size == 0 {
		return nil
	}
	if node == nil {
		return nil
	} else if node.key == key {
		return node.value
	}

	if key.(Comparable).Compare(node.key) < 0 {
		return tm.get(node.left, key)
	} else {
		return tm.get(node.right, key)
	}
}
func (tm *TreeMap) Set(key interface{}, newValue interface{}) {
	tm.set(tm.root, key, newValue)
}
func (tm *TreeMap) set(node *treeMapNode, key interface{}, newValue interface{}) {
	if tm.size == 0 {
		panic("bu cun zai zhe yang de key")
	}
	if node == nil {
		panic("bu cun zai zhe yang de key1")
	}
	if node.key == key {
		node.value = newValue
		return
	}

	if key.(Comparable).Compare(node.key) < 0 {
		tm.set(node.left, key, newValue)
	} else {
		tm.set(node.right, key, newValue)
	}
}
func (tm *TreeMap) GetSize() int {
	return tm.size
}

func (tm *TreeMap) IsEmpty() bool {
	return tm.size == 0
}

func (tm *TreeMap) String() string {
	str := ""
	return createString2(tm.root, 0, str)
}

func createString2(node *treeMapNode, depth int, str string) string {
	if node == nil {
		str += depthString2(depth) + "nil \n"
		return str
	}

	str += depthString2(depth) + "key:" + fmt.Sprint(node.key) + " value:" + fmt.Sprint(node.value) + "\n"
	str = createString2(node.left, depth+1, str)
	str = createString2(node.right, depth+1, str)
	return str
}
func depthString2(depth int) string {
	str := fmt.Sprint("")
	for i := 0; i < depth; i++ {
		str += fmt.Sprint("--")
	}
	return str
}
