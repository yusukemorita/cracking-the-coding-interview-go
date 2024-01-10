package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	removeDuplicates := map[string](func(node *LinkedListNode)){
		"RemoveDuplicates1": RemoveDuplicates1,
		"RemoveDuplicates2": RemoveDuplicates2,
	}

	for name, removeDuplicatesFunc := range removeDuplicates {
		t.Run(name, func(t *testing.T) {

			testCases := map[string]struct {
				linkedList *LinkedListNode
				expected   []string
			}{
				"returns original list when no duplicates": {
					linkedList: &LinkedListNode{
						value: "hi",
						pointer: &LinkedListNode{
							value: "hallo",
						},
					},
					expected: []string{"hi", "hallo"},
				},
				"removes duplicates when duplicates are present": {
					linkedList: &LinkedListNode{
						value: "hi",
						pointer: &LinkedListNode{
							value: "hello",
							pointer: &LinkedListNode{
								value: "hi",
							},
						},
					},
					expected: []string{"hi", "hello"},
				},
				"removes duplicates when duplicates are present (little more complex)": {
					linkedList: &LinkedListNode{
						value: "hi",
						pointer: &LinkedListNode{
							value: "hello",
							pointer: &LinkedListNode{
								value: "hi",
								pointer: &LinkedListNode{
									value: "hallo",
									pointer: &LinkedListNode{
										value: "hello",
										pointer: &LinkedListNode{
											value: "bonjour",
											pointer: &LinkedListNode{
												value: "konnichiwa",
											},
										},
									},
								},
							},
						},
					},
					expected: []string{"hi", "hello", "hallo", "bonjour", "konnichiwa"},
				},
			}

			for name, testCase := range testCases {
				t.Run(name, func(t *testing.T) {
					removeDuplicatesFunc(testCase.linkedList)
					got := testCase.linkedList.values()
					if !reflect.DeepEqual(got, testCase.expected) {
						t.Errorf("expected %v, got %v", testCase.expected, got)
					}
				})
			}
		})
	}
}
