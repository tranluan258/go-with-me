package tree

type BST struct {
	tree   *BinaryNode
	length int
}

func (t *BST) insert(val int) {
	if t.tree == nil {
		t.tree = &BinaryNode{
			val:   val,
			left:  nil,
			right: nil,
		}
		return
	}

	curr := t.tree

	for curr != nil {
		if val > curr.val {
			if curr.right == nil {
				curr.right = &BinaryNode{val: val, left: nil, right: nil}
				break
			}

			curr = curr.right
		}

		if val < curr.val {
			if curr.left == nil {
				curr.left = &BinaryNode{val: val, left: nil, right: nil}
				break
			}
			curr = curr.left
		}

	}

	t.length++
}
