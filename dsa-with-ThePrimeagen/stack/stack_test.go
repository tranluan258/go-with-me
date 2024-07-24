package stack

import "testing"

func TestStack(t *testing.T) {
	newStack := &Stack{head: nil, length: 0}

	newStack.Push(1)

	if newStack.Size() != 1 {
		t.Fatalf("Failed test case expected %d rec %d", 1, newStack.Size())
	}

	newStack.Push(2)

	if newStack.Peek() != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, newStack.Peek())
	}

	newStack.Push(3)

	if newStack.Pop() != 3 {
		t.Fatalf("Failed test case expected %d rec %d", 3, newStack.Pop())
	}

	if newStack.Size() != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, newStack.Size())
	}

	if newStack.Peek() != 2 {
		t.Fatalf("Failed test case expected %d rec %d", 2, newStack.Peek())
	}
}
