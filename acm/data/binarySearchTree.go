package data

import (
	"fmt"
)

type BinarySearchTree struct {
	root *treeNode
	size int
}

type treeNode struct {
	value       Comparable
	left, right *treeNode
}

func CreateBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
	}
}

func (bst *BinarySearchTree) GetSize() int {
	return bst.size
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BinarySearchTree) Add(value Comparable) {
	bst.root = bst.add(bst.root, value)
}
func (bst *BinarySearchTree) add(node *treeNode, value Comparable) *treeNode {
	if node == nil {
		bst.size++
		return &treeNode{value, nil, nil}
	}

	if value.Compare(node.value) < 0 {
		node.left = bst.add(node.left, value)
	} else if value.Compare(node.value) > 0 {
		node.right = bst.add(node.right, value)
	}
	return node
}
func (bst *BinarySearchTree) Contains(value Comparable) bool {
	return contains(bst.root, value)
}
func contains(node *treeNode, value Comparable) bool {
	if node == nil {
		return false
	} else if node.value == value {
		return true
	}
	if value.Compare(node.value) < 0 {
		return contains(node.left, value)
	} else {
		return contains(node.right, value)
	}
}

func (bst *BinarySearchTree) PreOrder() { //前序遍历
	preOrder(bst.root)
}
func preOrder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	preOrder(node.left)
	preOrder(node.right)
}

func (bst *BinarySearchTree) MidOrder() { //中序遍历
	midOrder(bst.root)
}
func midOrder(node *treeNode) {
	if node == nil {
		return
	}
	midOrder(node.left)
	fmt.Println(node.value)
	midOrder(node.right)
}
func (bst *BinarySearchTree) PostOrder() { //中序遍历
	postOrder(bst.root)
}
func postOrder(node *treeNode) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.value)
}
func (bst *BinarySearchTree) LevelOrder() {
	lq := CreateLoopQueue(10)
	lq.EnQueue(bst.root)
	for !lq.IsEmpty() {
		cur := lq.DeQueue().(*treeNode)
		fmt.Println(cur.value)
		if cur.left != nil {
			lq.EnQueue(cur.left)
		}
		if cur.right != nil {
			lq.EnQueue(cur.right)
		}
	}
}
func (bst *BinarySearchTree) Min() Comparable { //找出最小值
	if bst.size == 0 {
		return nil
	}
	return min(bst.root).value
}
func min(node *treeNode) *treeNode {
	if node.left == nil {
		return node
	}
	return min(node.left)
}

func (bst *BinarySearchTree) RemoveMin() Comparable {
	ret := min(bst.root)
	bst.root = removeMin(bst.root)
	bst.size--
	return ret.value
}
func removeMin(node *treeNode) *treeNode {
	if node.left == nil {
		right := node.right
		node.right = nil
		return right
	}

	node.left = removeMin(node.left)
	return node
}
func (bst *BinarySearchTree) Max() Comparable { //找出最大值
	if bst.size == 0 {
		return nil
	}
	return max(bst.root).value
}
func max(node *treeNode) *treeNode {
	if node.right == nil {
		return node
	}
	return max(node.right)
}
func (bst *BinarySearchTree) RemoveMax() Comparable {
	ret := max(bst.root)
	bst.root = removeMax(bst.root)
	bst.size--
	return ret.value
}
func removeMax(node *treeNode) *treeNode {
	if node.right == nil {
		left := node.left
		node.right = nil
		return left
	}

	node.right = removeMax(node.right)
	return node
}

func (bst *BinarySearchTree) Remove(value Comparable) {
	bst.root = bst.remove(bst.root, value)
}
func (bst *BinarySearchTree) remove(node *treeNode, value Comparable) *treeNode {
	if node == nil {
		return nil
	}
	if value.Compare(node.value) < 0 {
		node.left = bst.remove(node.left, value)
		return node
	} else if value.Compare(node.value) > 0 {
		node.right = bst.remove(node.right, value)
		return node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			bst.size--
			return right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			bst.size--
			return left
		}

		ret := min(node.right)
		ret.right = removeMin(node.right)
		ret.left = node.left
		node.left, node.right = nil, nil
		return ret

	}

}
func (bst *BinarySearchTree) String() string {
	str := ""
	return createString(bst.root, 0, str)
}
func createString(node *treeNode, depth int, str string) string {
	if node == nil {
		str += DepthString(depth) + "nil \n"
		return str
	}

	str += DepthString(depth) + fmt.Sprint(node.value) + "\n"
	str = createString(node.left, depth+1, str)
	str = createString(node.right, depth+1, str)
	return str
}
func DepthString(depth int) string {
	str := fmt.Sprint("")
	for i := 0; i < depth; i++ {
		str += fmt.Sprint("--")
	}
	return str
}
