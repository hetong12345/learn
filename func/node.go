package main

type node struct {
	value int
	left  *node
	right *node
}

func (node *node) trav(f func(*node)) {
	if node == nil {
		return
	}
	node.left.trav(f)
	node.trav(f)
	node.right.trav(f)
}
