package stack

type StackNode struct {
	next  *StackNode
	value int
}

type Stack struct {
	head   *StackNode
	length int
}

func (q *Stack) Pop() int {
	if q.head == nil {
		return -1
	}

	q.length--
	head := q.head
	q.head = head.next
	head.next = nil
	return head.value
}

func (q *Stack) Push(val int) {
	node := &StackNode{value: val, next: nil}

	if q.head == nil {
		q.head = node
		q.length++
		return
	}

	head := q.head
	q.head = node
	q.head.next = head
	q.length++
}

func (q *Stack) Peek() int {
	if q.head == nil {
		return -1
	}
	return q.head.value
}

func (q *Stack) Size() int {
	return q.length
}
