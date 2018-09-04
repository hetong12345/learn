package data

type SegmentTree struct {
	data   []interface{}
	tree   []interface{}
	merger Merger
}

type Merger func(interface{}, interface{}) interface{}

func CreateSegmentTree(arr []interface{}, merger Merger) *SegmentTree {
	tree := make([]interface{}, len(arr)*4)
	st := &SegmentTree{
		data:   arr,
		tree:   tree,
		merger: merger,
	}
	st.buildSegmentTree(0, 0, len(arr)-1)
	return st
}
func (st *SegmentTree) buildSegmentTree(index, l, r int) {
	if l == r {
		st.tree[index] = st.data[l]
		return
	}

	left := st.leftChild(index)
	right := st.rightChild(index)

	mid := l + (r-l)/2

	st.buildSegmentTree(left, l, mid)
	st.buildSegmentTree(right, mid+1, r)

	st.tree[index] = st.merger(st.tree[left], st.tree[right])
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) leftChild(index int) int {
	return index*2 + 1
}

func (st *SegmentTree) rightChild(index int) int {
	return index*2 + 2
}
