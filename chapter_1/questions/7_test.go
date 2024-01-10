package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	testCases := map[string]struct {
		input    [][]int
		expected [][]int
	}{
		"1": {
			input: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		"2": {
			input: [][]int{
				{0, 0, 1, 0},
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			fmt.Println("input")
			printMatrix(tc.input)

			output := RotateMatrix(tc.input, len(tc.input))

			fmt.Println("output")
			printMatrix(output)

			if len(tc.expected) != len(output) {
				t.Errorf("expected length %d, got length %d", len(tc.expected), len(output))
			}

			for index, expectedRow := range tc.expected {
				if !reflect.DeepEqual(expectedRow, output[index]) {
					t.Errorf("row %d does not match", index)
				}
			}
		})
	}
}

func printMatrix(matrix [][]int) {
	fmt.Println("")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("")
}
