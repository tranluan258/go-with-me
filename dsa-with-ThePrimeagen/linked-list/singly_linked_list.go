package linkedlist

type node struct {
	next *node
	val  int
}

type linkedlist struct {
	head   *node
	length int
}

func NewSinglyLinkedList() *linkedlist {
	return &linkedlist{
		head:   nil,
		length: 0,
	}
}

func (l *linkedlist) Find(val int) *node {
	if l.head == nil {
		return nil
	}

	current := l.head
	for current.next != nil {
		if current.val == val {
			return current
		}
		current = current.next
	}

	if current.val == val {
		return current
	}

	return nil
}

func (l *linkedlist) Insert(val int) {
	node := &node{val: val, next: nil}
	if l.head == nil {
		l.head = node
		l.length = 1
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
		l.length++
	}
}

func (l *linkedlist) Head() *node {
	return l.head
}

func (l *linkedlist) Size() int {
	return l.length
}

func (l *linkedlist) Last() *node {
	if l.head == nil {
		return nil
	}
	curr := l.head

	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

func (l *linkedlist) Remove(val int) {
	if l.head == nil {
		return
	}

	if l.head.val == val {
		l.head = l.head.next
		l.length--
		return
	}

	curr := l.head
	var prev *node

	for curr.next != nil {
		if curr.val == val {
			prev.next = curr.next
			l.length--
			break
		}
		prev = curr
		curr = curr.next
	}
}

func (l *linkedlist) Reverse() {
	var prev *node = nil
	var next *node = nil
	curr := l.head

	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}
	l.head = prev
}
