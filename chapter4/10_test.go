package main

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	t.Run("returns true when subtree", func(t *testing.T) {
		//     10
		//   5    15
		//       13 20
		tree := &BinaryTreeNode{value: 10}
		left := &BinaryTreeNode{
			value:  5,
			parent: tree,
		}
		tree.left = left
		right := &BinaryTreeNode{
			value:  15,
			parent: tree,
		}
		tree.right = right
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

		//   15
		//  13 20
		subtree := &BinaryTreeNode{
			value:  15,
			parent: tree,
		}
		subtreeRight := &BinaryTreeNode{
			value:  20,
			parent: subtree,
		}
		subtree.right = subtreeRight
		subtreeLeft := &BinaryTreeNode{
			value:  13,
			parent: subtree,
		}
		subtree.left = subtreeLeft

		got := IsSubtree(tree, subtree)
		if !got {
			t.Error()
		}
	})

	t.Run("returns false when not subtree", func(t *testing.T) {
		//     10
		//   5    15
		//       13 20
		tree := &BinaryTreeNode{value: 10}
		left := &BinaryTreeNode{
			value:  5,
			parent: tree,
		}
		tree.left = left
		right := &BinaryTreeNode{
			value:  15,
			parent: tree,
		}
		tree.right = right
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

		//   15
		//  13 20
		subtree := &BinaryTreeNode{
			value:  15,
			parent: tree,
		}
		subtreeRight := &BinaryTreeNode{
			value:  20,
			parent: subtree,
		}
		subtree.right = subtreeRight
		subtreeLeft := &BinaryTreeNode{
			value:  14, // different
			parent: subtree,
		}
		subtree.left = subtreeLeft

		got := IsSubtree(tree, subtree)
		if got {
			t.Error()
		}
	})
}
