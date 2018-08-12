package data

import "fmt"

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
	bst.root = add(bst.root, value)
	bst.size++
}

func add(node *treeNode, value Comparable) *treeNode {
	if node == nil {
		return &treeNode{value, nil, nil}
	}

	if value.CompareTo(node.value) < 0 {
		node.left = add(node.left, value)
	} else if value.CompareTo(node.value) > 0 {
		node.right = add(node.right, value)
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
	if value.CompareTo(node.value) < 0 {
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

func (bst *BinarySearchTree) RemoveMin() Comparable { //todo 完成删除最大最小的递归操作
	ret := min(bst.root)
	//right:=ret.right
	//removeMin(ret)
	return nil
}
func removeMin(node *treeNode) Comparable {
	//if
	return nil
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
	ret := bst.Max()
	return ret
}
func (bst *BinarySearchTree) String() string {
	str := ""
	return createString(bst.root, 0, str)
}
func createString(node *treeNode, depth int, str string) string {
	if node == nil {
		str += deepthString(depth) + "nil \n"
		return str
	}

	str += deepthString(depth) + fmt.Sprint(node.value) + "\n"
	str = createString(node.left, depth+1, str)
	str = createString(node.right, depth+1, str)
	return str
}
func deepthString(depth int) string {
	str := fmt.Sprint("")
	for i := 0; i < depth; i++ {
		str += fmt.Sprint("--")
	}
	return str
}
