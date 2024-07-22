package queue

type QNode struct {
	next  *QNode
	value int
}

type Queue struct {
	head   *QNode
	tail   *QNode
	length int
}

func (q *Queue) Dequeue() int {
	if q.head == nil {
		return -1
	}

	head := q.head
	q.head = head.next
	q.length--
	head.next = nil
	return head.value
}

func (q *Queue) Enqueue(val int) {
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

func (q *Queue) Peek() int {
	if q.head == nil {
		return -1
	}
	return q.head.value
}

func (q *Queue) Size() int {
	return q.length
}
