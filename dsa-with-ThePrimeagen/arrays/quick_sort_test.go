package arrays

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input, sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 6, 4, 3, 5, 7, 10, 1, 11, 9}, []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11}},
	}
	for i, test := range tests {
		QuickSort(test.input)
		if !slices.Equal(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}
