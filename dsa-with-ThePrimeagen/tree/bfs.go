package tree

func BFS(head *BinaryNode, needle int) bool {
	q := &BTQueue{
		head: &QNode{
			value: head,
			next:  nil,
		},
		tail:   nil,
		length: 1,
	}

	for q.length > 0 {
		curr := q.Dequeue()
		if curr == nil {
			return false
		}

		if curr.val == needle {
			return true
		}

		if curr.left != nil {
			q.Enqueue(curr.left)
		}

		if curr.right != nil {
			q.Enqueue(curr.right)
		}
	}

	return false
}
