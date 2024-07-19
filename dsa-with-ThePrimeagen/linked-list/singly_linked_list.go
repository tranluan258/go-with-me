package linkedlist

type Node[T int] struct {
	next *Node[T]
	key  T
}

type LinkedList[T int] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Find(key T) *Node[T] {
	if l.head == nil {
		return nil
	}

	current := l.head
	for current.next != nil {
		if current.key == key {
			return current
		}
		current = current.next
	}

	if current.key == key {
		return current
	}

	return nil
}

func (l *LinkedList[T]) Insert(key T) {
	node := &Node[T]{key: key, next: nil}
	if l.head == nil {
		l.head = node
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) Last() *Node[T] {
	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}
	return curr
}
