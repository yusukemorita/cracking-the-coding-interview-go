package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	//   15
	//  13 20
	root := &BinaryTreeNode{value:  15}
	right := &BinaryTreeNode{value:  20}
	root.right = right
	left := &BinaryTreeNode{value:  13}
	root.left = left

	tree := BinaryTree{root: root}

	if right.right != nil {
		t.Fatal("right.right must be nil")
	}

	tree.Insert(25)

	if right.right == nil {
		t.Fatal("right.right should not be nil")
	}
	if right.right.value != 25 {
		t.Error("right.right should be 25")
	}
}
