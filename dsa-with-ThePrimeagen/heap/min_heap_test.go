package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	minHeap := &MinHeap{
		data:   make([]int, 100),
		length: 0,
	}

	minHeap.insert(1)
	minHeap.insert(3)
	minHeap.insert(2)
	minHeap.insert(4)
	minHeap.insert(5)

	if minHeap.length != 5 {
		t.Fatalf("expected %d rec %d", 5, minHeap.length)
	}

	out := minHeap.delete()

	if out != 1 {
		t.Fatalf("expected %d rec %d", 1, out)
	}

	if minHeap.data[0] != 2 {
		t.Fatalf("expected %d rec %d", 2, minHeap.data[0])
	}
	minHeap.delete()
	minHeap.delete()
	minHeap.delete()
	endOunt := minHeap.delete()

	if endOunt != 5 || minHeap.length != 0 {
		t.Fatalf("expected %d rec %d", 1, out)
	}
}
