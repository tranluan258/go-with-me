package tree

import "testing"

func TestBst(t *testing.T) {
	tree := &BST{
		tree:   nil,
		length: 0,
	}

	tree.insert(2)
	tree.insert(1)
	tree.insert(3)
	if tree.tree.val != 2 {
		t.Fatalf("expected %d rec %d", 2, tree.tree.val)
	}

	if tree.tree.right.val != 3 {
		t.Fatalf("expected %d rec %d", 3, tree.tree.right.val)
	}
	if tree.tree.left.val != 1 {
		t.Fatalf("expected %d rec %d", 1, tree.tree.left.val)
	}
}
