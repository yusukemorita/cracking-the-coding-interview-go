package main

type BinaryTree struct {
	root *BinaryTreeNode
}

func (tree *BinaryTree) Insert(value int) {
	insert(tree.root, value)
}

func insert(root *BinaryTreeNode, value int) {
	if root.value >= value {
		if root.left == nil {
			root.left = &BinaryTreeNode{value: value}
		} else {
			insert(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &BinaryTreeNode{value: value}
		} else {
			insert(root.right, value)
		}
	}
}

// func (tree *BinaryTree) Find(val int) *BinaryTreeNode {

// }

// func (tree *BinaryTree) Delete(node *BinaryTreeNode) {

// }

// func (tree *BinaryTree) GetRandomNode() *BinaryTreeNode {

// }
