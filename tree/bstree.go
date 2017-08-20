package tree

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type BSTree struct {
	root *node
	size int
}

func (t *BSTree) Add(data int) {
	t.root = add(data, t.root)
	t.size++
}

func (t *BSTree) InOrder() {
	inOrder(t.root)
}

func inOrder(n *node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	fmt.Print(n.data)
	inOrder(n.right)
}

func add(data int, n *node) *node {
	if n == nil {
		return &node{data: data, left: nil, right: nil}
	}

	if data >= n.data {
		n.right = add(data, n.right)
	} else {
		n.left = add(data, n.left)
	}

	return n
}
