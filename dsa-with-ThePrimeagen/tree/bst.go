package tree

type BST struct {
	tree   *BinaryNode
	length int
}

func (t *BST) insert(val int) {
	t.length++
	if t.tree == nil {
		t.tree = &BinaryNode{
			val:   val,
			left:  nil,
			right: nil,
		}
		return
	}

	curr := t.tree

forLabel:
	for curr != nil {
		if val > curr.val {
			if curr.right == nil {
				curr.right = &BinaryNode{val: val, left: nil, right: nil}
				break forLabel
			}

			curr = curr.right
		} else if val < curr.val {
			if curr.left == nil {
				curr.left = &BinaryNode{val: val, left: nil, right: nil}
				break forLabel
			}
			curr = curr.left
		}
	}
}

func (t *BST) find(val int) bool {
	curr := t.tree

	for curr != nil {
		if curr.val == val {
			return true
		}

		if val > curr.val {
			curr = curr.right
		} else if val < curr.val {
			curr = curr.left
		}

	}
	return false
}
