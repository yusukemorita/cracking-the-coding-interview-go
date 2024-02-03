package main

import (
	"fmt"
	"testing"
)

func TestCommonAncestor(t *testing.T) {
	//     10
	//   5    15
	//       13 20
	root := &BinaryTreeNode{value: 10}
	left := &BinaryTreeNode{
		value:  5,
		parent: root,
	}
	root.left = left
	right := &BinaryTreeNode{
		value:  15,
		parent: root,
	}
	root.right = right
	rightRight := &BinaryTreeNode{
		value:  20,
		parent: right,
	}
	right.right = rightRight
	rightLeft := &BinaryTreeNode{
		value:  13,
		parent: right,
	}
	right.left = rightLeft

	got := CommonAncestor(root, right)
	if got != root {
		t.Error()
	}

	got = CommonAncestor(left, right)
	if got != root {
		t.Error()
	}

	got = CommonAncestor(left, rightRight)
	if got != root {
		t.Error()
	}

	got = CommonAncestor(rightRight, rightLeft)
	if got != right {
		t.Error()
	}

	got = CommonAncestor(rightRight, right)
	if got != right {
		t.Error()
	}
}

func TestCommonAncestor2(t *testing.T) {
	//     10
	//   5    15
	//       13 20
	root := &BinaryTreeNode{value: 10}
	left := &BinaryTreeNode{
		value:  5,
		parent: root,
	}
	root.left = left
	right := &BinaryTreeNode{
		value:  15,
		parent: root,
	}
	root.right = right
	rightRight := &BinaryTreeNode{
		value:  20,
		parent: right,
	}
	right.right = rightRight
	rightLeft := &BinaryTreeNode{
		value:  13,
		parent: right,
	}
	right.left = rightLeft

	testCases := []struct {
		nodeA    *BinaryTreeNode
		nodeB    *BinaryTreeNode
		expected *BinaryTreeNode
	}{
		{
			nodeA:    root,
			nodeB:    right,
			expected: root,
		},
		{
			nodeA:    left,
			nodeB:    right,
			expected: root,
		},
		{
			nodeA:    left,
			nodeB:    rightRight,
			expected: root,
		},
		{
			nodeA:    rightLeft,
			nodeB:    rightRight,
			expected: right,
		},
		{
			nodeA:    rightLeft,
			nodeB:    right,
			expected: right,
		},
	}

	for index, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", index), func(t *testing.T) {
			got := CommonAncestor2(tc.nodeA, tc.nodeB)
			if got != tc.expected {
				t.Error()
			}
		})
	}
}
