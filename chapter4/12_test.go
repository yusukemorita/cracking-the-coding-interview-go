package main

import (
	"testing"
)

func TestCountPathsWithSum2(t *testing.T) {
	t.Run("when no negative numbers", func(t *testing.T) {
		//   15
		//  15 20
		//       10
		root := &BinaryTreeNode{value: 15}
		root.right = &BinaryTreeNode{value: 20}
		root.left = &BinaryTreeNode{value: 15}
		root.right.right = &BinaryTreeNode{value: 10}

		count := CountPathsWithSum2(root, 35)
		if count != 1 {
			t.Errorf("expected 1, got %d", count)
		}

		count = CountPathsWithSum2(root, 30)
		if count != 2 {
			t.Errorf("expected 2, got %d", count)
		}
	})

	t.Run("when negative numbers", func(t *testing.T) {
		//   15
		//  20 -15
		//       20
		root := &BinaryTreeNode{value: 15}
		root.right = &BinaryTreeNode{value: -15}
		root.left = &BinaryTreeNode{value: 20}
		root.right.right = &BinaryTreeNode{value: 20}

		count := CountPathsWithSum2(root, 20)
		if count != 3 {
			t.Errorf("expected 3, got %d", count)
		}
	})
}
