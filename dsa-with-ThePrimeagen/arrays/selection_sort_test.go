package arrays

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input, sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{4, 4, 4, 3, 5}, []int{3, 4, 4, 4, 5}},
	}
	for i, test := range tests {
		SelectionSort(test.input)
		if !slices.Equal(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}
