package tree

type QNode struct {
	next  *QNode
	value *BinaryNode
}

type BTQueue struct {
	head   *QNode
	tail   *QNode
	length int
}

func (q *BTQueue) Dequeue() *BinaryNode {
	if q.head == nil {
		return nil
	}

	head := q.head
	q.head = head.next
	q.length--
	head.next = nil
	return head.value
}

func (q *BTQueue) Enqueue(val *BinaryNode) {
	node := &QNode{value: val, next: nil}
	q.length++
	if q.tail == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *BTQueue) Peek() *BinaryNode {
	if q.head == nil {
		return nil
	}
	return q.head.value
}

func (q *BTQueue) Size() int {
	return q.length
}
