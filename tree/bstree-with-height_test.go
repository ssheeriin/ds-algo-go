package tree

import (
"testing"
)

func TestBSTreeWithHeight_Add(t *testing.T) {
	bst := new(tree)
	bst.Add(1)
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(9)
	bst.Add(2)
	bst.Add(6)
	bst.Add(4)
	bst.Add(8)

	height:= bst.Height()
	expected := 5
	if expected != height {
		t.Errorf("expected: %s, actual:%s", expected, height)
	}
}


