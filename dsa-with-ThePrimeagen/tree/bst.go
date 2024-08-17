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

func (t *BST) delete(val int) {
	t.length--
	curr := t.tree
	var parent *BinaryNode

	for curr != nil {
		if curr.val == val {
			if curr.left == nil && curr.right == nil {
				t.removeNodeDontHaveLeftRight(curr, parent)
				break
			} else if curr.left != nil && curr.right != nil {
				t.removeNodeHaveBothLeftRight(curr, parent)
				break
			} else {
				t.removeNodeOnlyHaveLeftOrRight(curr, parent)
				break
			}
		}

		parent = curr
		if val > curr.val {
			curr = curr.right
		} else if val < curr.val {
			curr = curr.left
		}

	}
}

func (t *BST) removeNodeDontHaveLeftRight(curr *BinaryNode, parent *BinaryNode) {
	if curr.val > parent.val {
		parent.right = nil
	} else {
		parent.left = nil
	}
	curr = nil
}

func (t BST) removeNodeHaveBothLeftRight(curr *BinaryNode, parent *BinaryNode) {
	tmp := curr
	var prev *BinaryNode

	if curr.val > parent.val {
		for curr.left != nil {
			prev = curr
			curr = curr.left
		}

		if prev != nil {
			prev.left = nil
		}

		parent.right = curr
		curr.right = tmp.right

	} else {

		for curr.right != nil {
			prev = curr
			curr = curr.right
		}

		if prev != nil {
			prev.left = nil
		}

		parent.left = curr
		curr.left = tmp.left
	}
	return
}

func (t *BST) removeNodeOnlyHaveLeftOrRight(curr *BinaryNode, parent *BinaryNode) {
	if curr.val > parent.val {
		if curr.left == nil {
			parent.right = curr.right
			curr = nil
		} else if curr.right == nil {
			parent.right = curr.left
			curr = nil
		}
		return
	}

	if curr.val < parent.val {
		if curr.left == nil {
			parent.left = curr.right
			curr = nil
		} else if curr.right == nil {
			parent.left = curr.left
			curr = nil
		}
		return
	}
}
