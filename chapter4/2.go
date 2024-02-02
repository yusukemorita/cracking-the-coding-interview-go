package main

func BuildBinarySearchTree(sortedArray []int) *BinaryTreeNode {
	if len(sortedArray) == 0 {
		return nil
	}

	middleIndex := len(sortedArray) / 2
	middle := sortedArray[middleIndex]
	left := sortedArray[:middleIndex]
	right := sortedArray[middleIndex+1:]

	return &BinaryTreeNode{
		value: middle,
		left:  BuildBinarySearchTree(left),
		right: BuildBinarySearchTree(right),
	}
}

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (node *BinaryTreeNode) Height() int {
	leftHeight := 0
	if node.left != nil {
		leftHeight = node.left.Height()
	}

	rightHeight := 0
	if node.right != nil {
		rightHeight = node.right.Height()
	}

	if leftHeight >= rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
