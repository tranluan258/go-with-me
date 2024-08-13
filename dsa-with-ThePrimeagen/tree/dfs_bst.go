package tree

func dsf(head *BinaryNode, needle int) bool {
	if head == nil {
		return false
	}

	if head.val == needle {
		return true
	}

	if head.val < needle {
		return dsf(head.right, needle)
	}
	return dsf(head.left, needle)
}
