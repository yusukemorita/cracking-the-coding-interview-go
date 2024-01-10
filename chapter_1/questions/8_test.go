package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	for name, zeroMatrixFunction := range map[string](func(matrix [][]int) [][]int){
		// "ZeroMatrix1": ZeroMatrix1,
		"ZeroMatrix2": ZeroMatrix2,
	} {
		t.Run(name, func(t *testing.T) {
			testCases := map[string]struct {
				input    [][]int
				expected [][]int
			}{
				"1": {
					input: [][]int{
						{0, 1},
						{1, 1},
					},
					expected: [][]int{
						{0, 0},
						{0, 1},
					},
				},
				"1.5": {
					input: [][]int{
						{1, 0},
						{1, 1},
					},
					expected: [][]int{
						{0, 0},
						{1, 0},
					},
				},
				"2": {
					input: [][]int{
						{0, 1, 1},
						{1, 0, 1},
						{1, 1, 1},
					},
					expected: [][]int{
						{0, 0, 0},
						{0, 0, 0},
						{0, 0, 1},
					},
				},
				"empty matrix": {
					input:    [][]int{},
					expected: [][]int{},
				},
			}

			for name, tc := range testCases {
				t.Run(name, func(t *testing.T) {
					fmt.Println("input")
					printMatrix(tc.input)

					output := zeroMatrixFunction(tc.input)

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
		})
	}
}
