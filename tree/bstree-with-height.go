package tree

type hnode struct {
	data   int
	height int
	left   *hnode
	right  *hnode
}

type tree struct {
	root *hnode
	size int
}

func (t *tree) Add(data int) {
	t.root = addh(data, t.root)
	t.size++
}

func (t *tree) Height() int {
	return t.root.height
}
func addh(data int, n *hnode) *hnode {

	if n == nil {
		return &hnode{data: data, left: nil, right: nil, height: 1}
	}

	if data >= n.data {
		added := addh(data, n.right)
		n.right = added
		n.height = added.height + 1
	} else {
		added := addh(data, n.left)
		n.left = added
		n.height = added.height + 1
	}

	return n
}


