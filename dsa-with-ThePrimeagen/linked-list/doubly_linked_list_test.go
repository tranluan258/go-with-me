package linkedlist

import "testing"

func TestDoublyLinkedlist(t *testing.T) {
	var doublyLinkedlist DoublyLinkedlist = &doublyLinkedlist{head: nil, tail: nil, length: 0}

	doublyLinkedlist.append(1)

	if doublyLinkedlist.get(0) != 1 {
		t.Fatalf("expected %d rec %d", 1, doublyLinkedlist.get(0))
	}

	doublyLinkedlist.append(2)

	if doublyLinkedlist.get(1) != 2 {
		t.Fatalf("expected %d rec %d", 2, doublyLinkedlist.get(1))
	}

	doublyLinkedlist.prepend(3)

	if doublyLinkedlist.get(0) != 3 {
		t.Fatalf("expected %d rec %d", 3, doublyLinkedlist.get(0))
	}

	doublyLinkedlist.insertAt(4, 2)

	if doublyLinkedlist.get(2) != 2 {
		t.Fatalf("expected %d rec %d", 2, doublyLinkedlist.get(2))
	}
	if doublyLinkedlist.get(3) != 4 {
		t.Fatalf("expected %d rec %d", 4, doublyLinkedlist.get(4))
	}

	removeAt := doublyLinkedlist.removeAt(5, 1)
	if removeAt != -1 {
		t.Fatalf("expected %d rec %d", -1, removeAt)
	}

	doublyLinkedlist.remove(4)
	removeHead := doublyLinkedlist.remove(3)
	if doublyLinkedlist.get(0) != 1 {
		t.Fatalf("expected %d rec %d", 1, removeHead)
	}

	doublyLinkedlist.remove(2)
	doublyLinkedlist.remove(1)
	if doublyLinkedlist.size() != 0 {
		t.Fatalf("expected %d rec %d", 0, doublyLinkedlist.size())
	}
}
