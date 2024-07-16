package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []int{1, 2, 3, 5, 6, 7, 8, 9}
	v := 9

	if !BinarySearch(tests, 0, len(tests), v) {
		t.Fatal("Failed test BinarySearch")
	}
}
