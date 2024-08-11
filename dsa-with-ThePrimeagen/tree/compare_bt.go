package tree

func compare(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.val != b.val {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}
