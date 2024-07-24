package stack

type StackNode struct {
	prev  *StackNode
	value int
}

type Stack struct {
	head   *StackNode
	length int
}

func (s *Stack) Pop() int {
	if s.head == nil {
		return -1
	}

	s.length--
	head := s.head
	s.head = head.prev
	head.prev = nil
	return head.value
}

func (s *Stack) Push(val int) {
	node := &StackNode{value: val, prev: nil}

	s.length++
	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack) Peek() int {
	if s.head == nil {
		return -1
	}
	return s.head.value
}

func (s *Stack) Size() int {
	return s.length
}
