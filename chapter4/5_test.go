package main

import "testing"

func TestIsBinarySearchTree(t *testing.T) {
	t.Run("returns true when left <= node <= right for all nodes", func(t *testing.T) {
		node := &BinaryTreeNode{
			value: 5,
			left: &BinaryTreeNode{
				value: 3,
				left: &BinaryTreeNode{
					value: 2,
				},
			},
			right: &BinaryTreeNode{
				value: 7,
				left: &BinaryTreeNode{
					value: 6,
				},
				right: &BinaryTreeNode{
					value: 8,
				},
			},
		}
		got := IsBinarySearchTree(node)
		if !got {
			t.Error()
		}
	})

	t.Run("returns false when left > node exists", func(t *testing.T) {
		node := &BinaryTreeNode{
			value: 5,
			left: &BinaryTreeNode{
				value: 3,
				left: &BinaryTreeNode{
					value: 4,
				},
			},
			right: &BinaryTreeNode{
				value: 7,
				left: &BinaryTreeNode{
					value: 6,
				},
				right: &BinaryTreeNode{
					value: 8,
				},
			},
		}
		got := IsBinarySearchTree(node)
		if got {
			t.Error()
		}
	})

	t.Run("returns false when node > right exists", func(t *testing.T) {
		node := &BinaryTreeNode{
			value: 5,
			left: &BinaryTreeNode{
				value: 3,
				left: &BinaryTreeNode{
					value: 2,
				},
			},
			right: &BinaryTreeNode{
				value: 9,
				left: &BinaryTreeNode{
					value: 6,
				},
				right: &BinaryTreeNode{
					value: 8,
				},
			},
		}
		got := IsBinarySearchTree(node)
		if got {
			t.Error()
		}
	})

	t.Run("returns false when left <= node.value <= right.value, but its not a BST when ancestors are inspected", func(t *testing.T) {
		node := &BinaryTreeNode{
			value: 20,
			left: &BinaryTreeNode{
				value: 10,
				right: &BinaryTreeNode{
					value: 25, // bigger than the root, 20, but on the left side of root
				},
			},
			right: &BinaryTreeNode{
				value: 30,
			},
		}
		got := IsBinarySearchTree(node)
		if got {
			t.Error()
		}
	})
}
