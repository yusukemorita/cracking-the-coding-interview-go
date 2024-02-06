package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeSort(t *testing.T) {
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
			got := MergeSort(tc.original)
			diff := cmp.Diff(got, tc.expected)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	half1 := []int{1, 3, 5}
	half2 := []int{2, 4}

	got := merge(half1, half2)

	diff := cmp.Diff(got, []int{1, 2, 3, 4, 5})
	if diff != "" {
		t.Error(diff)
	}
}
