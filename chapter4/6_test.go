package main

import "testing"

func TestNextNode(t *testing.T) {
	t.Run("returns next node when parent is the next node", func(t *testing.T) {
		//     10
		//   5    15
    //       13 20
		root := &BinaryTreeNode{
			value: 10,
		}

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

		next := NextNode(rightLeft)
		if next != right {
			t.Error()
		}
	})

	t.Run("returns next node when child is the next node", func(t *testing.T) {
		//     10
		//   5    15
    //       13 20
		root := &BinaryTreeNode{
			value: 10,
		}

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

		next := NextNode(right)
		if next != rightRight {
			t.Error()
		}
	})

	t.Run("returns next node when grandchild is the next node", func(t *testing.T) {
		//     10
		//   5    15
    //       13 20
		root := &BinaryTreeNode{
			value: 10,
		}

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

		next := NextNode(root)
		if next != rightLeft {
			t.Error()
		}
	})

	t.Run("returns nil when node has the largest value in the tree", func(t *testing.T) {
		//     10
		//   5    15
    //       13 20
		root := &BinaryTreeNode{
			value: 10,
		}

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

		next := NextNode(rightRight)
		if next != nil {
			t.Error()
		}
	})
}
