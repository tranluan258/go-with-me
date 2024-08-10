package tree

import "testing"

func TestBTQueue(t *testing.T) {
	newQueue := &BTQueue{head: nil, tail: nil, length: 0}

	newQueue.Enqueue(&BinaryNode{val: 1})

	if newQueue.Size() != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, newQueue.Size())
	}

	newQueue.Enqueue(&BinaryNode{val: 2})

	tail := newQueue.Peek()
	if tail == nil || tail.val != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, tail.val)
	}

	newQueue.Enqueue(&BinaryNode{val: 3})

	if newQueue.tail.value.val != 3 {
		t.Fatalf("Failed test case expected %d rec %d", 3, newQueue.tail.value.val)
	}

	val := newQueue.Dequeue()
	if val == nil || val.val != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, val.val)
	}

	if newQueue.Size() != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, newQueue.Size())
	}

	tail = newQueue.Peek()
	if tail == nil || tail.val != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, tail.val)
	}
}
