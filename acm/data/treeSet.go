package data

type TreeSet struct {
	bst *BinarySearchTree
}

func CreateTreeSet() *TreeSet {
	return &TreeSet{CreateBinarySearchTree()}
}
func (ts *TreeSet) GetSize() int {
	return ts.bst.GetSize()
}

func (ts *TreeSet) IsEmpty() bool {
	return ts.bst.IsEmpty()
}

func (ts *TreeSet) Add(value Comparable) {
	ts.bst.Add(value)
}

func (ts *TreeSet) Remove(value Comparable) {
	ts.bst.Remove(value)
}

func (ts *TreeSet) Contains(value Comparable) bool {
	return ts.bst.Contains(value)
}
func (ts *TreeSet) String() string {
	return ts.bst.String()
}
