package heap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	minHeap := &MaxHeap{
		data:   make([]int, 100),
		length: 0,
	}

	minHeap.insert(3)
	minHeap.insert(1)
	minHeap.insert(2)
	minHeap.insert(4)
	minHeap.insert(5)

	if minHeap.length != 5 {
		t.Fatalf("expected %d rec %d", 5, minHeap.length)
	}

	out := minHeap.delete()

	if out != 5 {
		t.Fatalf("expected %d rec %d", 5, out)
	}

	if minHeap.data[0] != 4 {
		t.Fatalf("expected %d rec %d", 4, minHeap.data[0])
	}
	minHeap.delete()
	minHeap.delete()
	minHeap.delete()
	endOunt := minHeap.delete()

	if endOunt != 1 || minHeap.length != 0 {
		t.Fatalf("expected %d rec %d", 1, out)
	}
}
