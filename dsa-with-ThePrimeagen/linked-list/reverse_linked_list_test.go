package linkedlist

import "testing"

func TestReverseLinkedlist(t *testing.T) {
	linkedlist := NewSinglyLinkedList()
	linkedlist.Insert(1)
	linkedlist.Insert(2)
	linkedlist.Insert(3)
	linkedlist.Insert(4)
	linkedlist.Insert(5)
	linkedlist.Insert(6)

	linkedlist.Reverse()

	if linkedlist.Head().val != 6 {
		t.Fatalf("Failed expected %d rec %d", 6, linkedlist.Head().val)
	}
}
