package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := map[string]struct {
		linkedList LinkedListNode
		expected   bool
	}{
		"returns true when is permutation": {
			linkedList: LinkedListNode{
				value: "hi",
				pointer: &LinkedListNode{
					value: "hello",
					pointer: &LinkedListNode{
						value: "hi",
					},
				},
			},
			expected: true,
		},
		"returns false when is not permutation": {
			linkedList: LinkedListNode{
				value: "hi",
				pointer: &LinkedListNode{
					value: "hello",
					pointer: &LinkedListNode{
						value: "bonjour",
					},
				},
			},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := IsPalindrome(&testCase.linkedList)
			if got != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	testCases := map[string]struct {
		linkedList LinkedListNode
		expected   bool
	}{
		"returns true when is permutation": {
			linkedList: LinkedListNode{
				value: "hi",
				pointer: &LinkedListNode{
					value: "hello",
					pointer: &LinkedListNode{
						value: "hi",
					},
				},
			},
			expected: true,
		},
		"returns false when is not permutation": {
			linkedList: LinkedListNode{
				value: "hi",
				pointer: &LinkedListNode{
					value: "hello",
					pointer: &LinkedListNode{
						value: "bonjour",
					},
				},
			},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := IsPalindrome2(&testCase.linkedList)
			if got != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := map[string]struct {
		linkedList *LinkedListNode
		expected   []string
	}{
		"returns reversed": {
			linkedList: &LinkedListNode{
				value: "a",
				pointer: &LinkedListNode{
					value: "b",
					pointer: &LinkedListNode{
						value: "c",
					},
				},
			},
			expected: []string{"c", "b", "a"},
		},
		"returns node when length is 1": {
			linkedList: &LinkedListNode{
				value: "a",
			},
			expected: []string{"a"},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got, _ := Reverse(testCase.linkedList)
			if !reflect.DeepEqual(got.values(), testCase.expected) {
				t.Errorf("expected %v, got %v", testCase.expected, got.values())
			}
		})
	}
}

func TestReverseWithPointers(t *testing.T) {
	testCases := map[string]struct {
		linkedList LinkedListNode
		expected   []string
	}{
		"returns reversed": {
			linkedList: LinkedListNode{
				value: "a",
				pointer: &LinkedListNode{
					value: "b",
					pointer: &LinkedListNode{
						value: "c",
					},
				},
			},
			expected: []string{"c", "b", "a"},
		},
		"returns node when length is 1": {
			linkedList: LinkedListNode{
				value: "a",
			},
			expected: []string{"a"},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := ReverseWithPointers(&testCase.linkedList)
			if !reflect.DeepEqual(got.values(), testCase.expected) {
				t.Errorf("expected %v, got %v", testCase.expected, got.values())
			}
		})
	}
}
