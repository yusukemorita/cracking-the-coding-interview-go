package main

import (
	"fmt"
	"testing"
)

// difficult to test, so just print the BST
func TestBuildDepthLists(t *testing.T) {
	t.Run("", func(t *testing.T) {
		bst := BinaryTreeNode{
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

		linkedLists := BuildDepthLists(&bst)
		if len(linkedLists) != 3 {
			t.Fatal("expected 3 lists")
		}

		fmt.Println(linkedLists[0].values())
		fmt.Println(linkedLists[1].values())
		fmt.Println(linkedLists[2].values())
	})
}
