package lintcode

func merge(i, j int) int { //lintcode 943  leetcode 303
	return i + j
}

type NumArray struct {
	data []int
	tree []int
}

func Constructor(nums []int) NumArray {
	if nums == nil || len(nums) == 0 {
		return NumArray{}
	}
	tree := make([]int, len(nums)*4)
	st := NumArray{
		data: nums,
		tree: tree,
	}
	st.buildSegmentTree(0, 0, len(nums)-1)
	return st
}
func (st *NumArray) buildSegmentTree(index, l, r int) {
	if l == r {
		st.tree[index] = st.data[l]
		return
	}

	left := st.leftChild(index)
	right := st.rightChild(index)
	mid := l + (r-l)/2

	st.buildSegmentTree(left, l, mid)
	st.buildSegmentTree(right, mid+1, r)

	st.tree[index] = merge(st.tree[left], st.tree[right])
}

func (st *NumArray) SumRange(qleft, qright int) int {
	if qleft < 0 || qleft >= len(st.data) || qright < 0 || qright >= len(st.data) || qleft > qright {
		panic("query condition is illegal")
	}

	return st.query(0, 0, len(st.data)-1, qleft, qright)
}

func (st *NumArray) query(index, l, r, qleft, qright int) int {
	if l == r {
		return st.tree[index]
	}

	left := st.leftChild(index)
	right := st.rightChild(index)
	mid := l + (r-l)/2

	if qleft >= mid+1 {
		return st.query(right, mid+1, r, qleft, qright)
	} else if qright <= mid {
		return st.query(left, l, mid, qleft, qright)
	}
	leftRes := st.query(left, l, mid, qleft, qright)
	rightRes := st.query(right, mid+1, r, qleft, qright)
	//return leftRes.merge(rightRes)
	return merge(leftRes, rightRes)
}
func (st *NumArray) Update(i int, val int) {
	if i < 0 || i >= len(st.data) {
		panic("i of tree is illegal")
	}
	st.data[i] = val
	st.set(0, 0, len(st.data)-1, i, val)
}

func (st *NumArray) set(tree, l, r, index int, e int) {
	if l == r {
		st.tree[tree] = e
		return
	}

	left := st.leftChild(tree)
	right := st.rightChild(tree)
	mid := l + (r-l)/2

	if index >= mid+1 {
		st.set(right, mid+1, r, index, e)
	} else {
		st.set(left, l, mid, index, e)
	}
	st.tree[tree] = merge(st.tree[left], st.tree[right])
}
func (st *NumArray) GetSize() int {
	return len(st.data)
}

func (st *NumArray) leftChild(index int) int {
	return index*2 + 1
}

func (st *NumArray) rightChild(index int) int {
	return index*2 + 2
}
