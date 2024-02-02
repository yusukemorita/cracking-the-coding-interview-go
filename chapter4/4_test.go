package main

import "testing"

func TestIsBalanced(t *testing.T) {
	t.Run("returns false when heights of subtrees differ by more than one", func(t *testing.T) {
		bst := &BinaryTreeNode{
			value: 6,
			left: &BinaryTreeNode{
				value: 4,
				left: &BinaryTreeNode{
					value: 2,
					right: &BinaryTreeNode{
						value: 5,
					},
				},
			},
			right: &BinaryTreeNode{
				value: 9,
			},
		}

		got, _ := IsBalanced(bst)
		if got {
			t.Error("expected false, got true")
		}
	})

	t.Run("returns true when heights of subtrees differ by one", func(t *testing.T) {
		bst := &BinaryTreeNode{
			value: 6,
			left: &BinaryTreeNode{
				value: 4,
				left: &BinaryTreeNode{
					value: 2,
					right: &BinaryTreeNode{
						value: 3,
					},
				},
				right: &BinaryTreeNode{
					value: 5,
				},
			},
			right: &BinaryTreeNode{
				value: 9,
				right: &BinaryTreeNode{
					value: 10,
				},
			},
		}
		got, _ := IsBalanced(bst)
		if !got {
			t.Error("expected true, got false")
		}
	})

	t.Run("returns false when heights of subtrees differ by zero", func(t *testing.T) {
		bst := &BinaryTreeNode{
			value: 6,
			left: &BinaryTreeNode{
				value: 4,
				left: &BinaryTreeNode{
					value: 2,
				},
			},
			right: &BinaryTreeNode{
				value: 9,
				right: &BinaryTreeNode{
					value: 10,
				},
			},
		}
		got, _ := IsBalanced(bst)
		if !got {
			t.Error("expected true, got false")
		}
	})
}
