package linkedlist

import "testing"

func TestSinglyLinkedlist(t *testing.T) {
	ll := NewSinglyLinkedList()
	ll.Insert(1)

	if ll.length != 1 {
		t.Fatalf("Failed test case insert Want %v got %v", 1, ll.length)
	}

	ll.Insert(4)
	ll.Insert(5)

	if ll.length != 3 {
		t.Fatalf("Failed test case insert Want %v got %v", 3, ll.length)
	}

	if ll.Head().val != 1 {
		t.Fatalf("Failed test case get Head Want %v got %v", 1, ll.Head().val)
	}

	if ll.Find(5) == nil {
		t.Fatalf("Failed test case find node Want %v got %v", 5, ll.Find(5).val)
	}

	if ll.Last().val != 5 {
		t.Fatalf("Failed test case last node Want %v got %v", 5, ll.Last().val)
	}

	ll.Remove(4)

	if ll.Find(4) != nil {
		t.Fatalf("Failed test case remove node Want %v got %v", nil, ll.Find(4))
	}

	ll.Insert(6)

	if ll.length != 3 {
		t.Fatalf("Failed test case insert Want %v got %v", 3, ll.length)
	}

	ll.Remove(1)

	if ll.Head().val != 5 || ll.length != 2 {
		t.Fatalf("Failed test case remove head Want %v got %v", 4, ll.Head().val)
	}
}
