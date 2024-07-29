package linkedlist

type dNode struct {
	prev *dNode
	next *dNode
	val  int
}

type DoublyLinkedlist interface {
	prepend(item int)
	append(item int)
	get(idx int) int
	insertAt(item int, idx int)
	remove(item int) int
	removeAt(item int, idx int) int
	size() int
}

type doublyLinkedlist struct {
	head   *dNode
	tail   *dNode
	length int
}

func (d *doublyLinkedlist) prepend(item int) {
	node := &dNode{val: item}
	d.length++
	if d.head == nil {
		d.head = node
		d.tail = node
		return
	}

	node.next = d.head
	d.head.prev = node
	d.head = node
}

func (d *doublyLinkedlist) append(item int) {
	node := &dNode{val: item}
	d.length++
	if d.tail == nil {
		d.head = node
		d.tail = node
		return
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node
}

func (d *doublyLinkedlist) insertAt(item, idx int) {
	if idx > d.length {
		return
	}

	switch idx {
	case 0:
		d.prepend(item)
		return
	case d.length - 1:
		d.append(item)
		return
	}

	curr := d.head
	for i := 0; i < idx && curr != nil; i++ {
		curr = curr.next
	}

	node := &dNode{val: item}
	d.length++
	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if curr.prev != nil {
		curr.prev.next = node
	}
}

func (d *doublyLinkedlist) remove(item int) int {
	curr := d.head
	for i := 0; i < d.length && curr != nil; i++ {
		if curr.val == item {
			break
		}
		curr = curr.next
	}
	node := d.removeNode(curr)

	if node == nil {
		return -1
	}
	return node.val
}

func (d *doublyLinkedlist) removeAt(item int, idx int) int {
	curr := d.head
	for i := 0; i < d.length && curr != nil; i++ {
		if curr.val == item {
			break
		}
		curr = curr.next
	}

	node := d.removeNode(curr)

	if node == nil {
		return -1
	}
	return node.val
}

func (d *doublyLinkedlist) removeNode(node *dNode) *dNode {
	if node == nil {
		return nil
	}
	d.length--

	if d.length == 0 {
		d.head = nil
		d.tail = nil
		return node
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node == d.head {
		d.head = node.next
	}

	if node == d.tail {
		d.tail = node.prev
	}

	node.next = nil
	node.prev = nil
	return node
}

func (d *doublyLinkedlist) size() int {
	return d.length
}

func (d *doublyLinkedlist) get(idx int) int {
	return d.getAt(idx)
}

func (d *doublyLinkedlist) getAt(idx int) int {
	curr := d.head
	for i := 0; i < idx && curr != nil; i++ {
		curr = curr.next
	}

	if curr == nil {
		return -1
	}

	return curr.val
}
