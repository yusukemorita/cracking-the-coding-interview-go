package main

import "testing"

func TestIsBalanced(t *testing.T) {
	t.Run("returns false when heights of subtrees differ by more than one", func(t *testing.T) {
		bst := &BSTNode{
			value: 6,
			left: &BSTNode{
				value: 4,
				left: &BSTNode{
					value: 2,
					right: &BSTNode{
						value: 5,
					},
				},
			},
			right: &BSTNode{
				value: 9,
			},
		}

		got, _ := IsBalanced(bst)
		if got {
			t.Error("expected false, got true")
		}
	})

	t.Run("returns true when heights of subtrees differ by one", func(t *testing.T) {
		bst := &BSTNode{
			value: 6,
			left: &BSTNode{
				value: 4,
				left: &BSTNode{
					value: 2,
					right: &BSTNode{
						value: 3,
					},
				},
				right: &BSTNode{
					value: 5,
				},
			},
			right: &BSTNode{
				value: 9,
				right: &BSTNode{
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
		bst := &BSTNode{
			value: 6,
			left: &BSTNode{
				value: 4,
				left: &BSTNode{
					value: 2,
				},
			},
			right: &BSTNode{
				value: 9,
				right: &BSTNode{
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
