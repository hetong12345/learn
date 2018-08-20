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
		return nil //todo complete this part
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
func (tm *TreeMap) Contains(key interface{}) bool {
	panic("implement me")
}

func (tm *TreeMap) Get(key interface{}) interface{} {
	panic("implement me")
}

func (tm *TreeMap) Set(key interface{}, newValue interface{}) {
	panic("implement me")
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
