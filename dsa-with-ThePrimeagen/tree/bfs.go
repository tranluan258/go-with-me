package tree

func BFS(head *BinaryNode, needle int) bool {
	q := &BTQueue{
		head: &QNode{
			value: head,
			next:  nil,
		},
		tail: &QNode{
			value: head,
			next:  nil,
		},
		length: 1,
	}

	for q.length > 0 {
		curr := q.Dequeue()
		if curr == nil {
			continue
		}

		if curr.val == needle {
			return true
		}

		if curr.left != nil {
			q.Enqueue(head.left)
		}

		if curr.right != nil {
			q.Enqueue(head.right)
		}
	}

	return false
}
