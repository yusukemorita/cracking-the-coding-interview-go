package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuickSort(t *testing.T) {
	testCases := map[string]struct {
		original []int
		expected []int
	}{
		"when slice is empty": {
			original: []int{},
			expected: []int{},
		},
		"when slice has one item": {
			original: []int{2},
			expected: []int{2},
		},
		"when slice is small": {
			original: []int{3, 1, 4, 2},
			expected: []int{1, 2, 3, 4},
		},
		"when slice is large": {
			original: []int{3, 8, 7, 1, 5, 4, 6, 9, 10, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			QuickSort(tc.original)
			diff := cmp.Diff(tc.expected, tc.original)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}
