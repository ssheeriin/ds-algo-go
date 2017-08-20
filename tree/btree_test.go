package tree

import "testing"

func TestBSTree_Add(t *testing.T) {
	bst := new(BSTree)
	bst.Add(1)
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(9)
	bst.Add(2)
	bst.Add(6)

	bst.InOrder()
}
