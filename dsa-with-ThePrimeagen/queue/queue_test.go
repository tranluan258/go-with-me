package queue

import "testing"

func TestQueue(t *testing.T) {
	newQueue := &Queue{head: nil, tail: nil, length: 0}

	newQueue.Enqueue(1)

	if newQueue.Size() != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, newQueue.Size())
	}

	newQueue.Enqueue(2)

	if newQueue.Peek() != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, newQueue.Peek())
	}

	newQueue.Enqueue(3)

	if newQueue.tail.value != 3 {
		t.Fatalf("Failed test case expected %d rec %d", 3, newQueue.tail.value)
	}

	if newQueue.Dequeue() != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, newQueue.Dequeue())
	}

	if newQueue.Size() != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, newQueue.Size())
	}

	if newQueue.Peek() != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, newQueue.Peek())
	}
}
