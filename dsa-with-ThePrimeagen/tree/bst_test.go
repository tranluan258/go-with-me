package tree

import "testing"

func TestBst(t *testing.T) {
	tree := &BST{
		tree:   nil,
		length: 0,
	}

	tree.insert(2)
	tree.insert(1)
	tree.insert(4)
	tree.insert(3)
	tree.insert(5)
	tree.insert(7)

	if !tree.find(3) {
		t.Fatalf("expected %t rec %t", true, false)
	}
	if !tree.find(5) {
		t.Fatalf("expected %t rec %t", true, false)
	}
	if tree.find(6) {
		t.Fatalf("expected %t rec %t", false, true)
	}

	tree.delete(4)
	if tree.find(4) {
		t.Fatalf("expected %t rec %t", false, true)
	}

	if tree.tree.right.val != 3 {
		t.Fatalf("expected %d rec %d", 3, tree.tree.right.val)
	}

	if !tree.find(5) {
		t.Fatalf("expected %t rec %t", true, false)
	}

	tree.delete(5)
	if tree.find(5) {
		t.Fatalf("expected %t rec %t", false, true)
	}

	tree.delete(7)
	if tree.find(7) {
		t.Fatalf("expected %t rec %t", false, true)
	}
}
