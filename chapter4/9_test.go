package main

import (
	"testing"
)

func TestBSTSequence(t *testing.T) {
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


	BSTSequence(root)
}
