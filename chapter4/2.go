package main

func BuildBinarySearchTree(sortedArray []int) *BSTNode {
	if len(sortedArray) == 0 {
		return nil
	}

	middleIndex := len(sortedArray) / 2
	middle := sortedArray[middleIndex]
	left := sortedArray[:middleIndex]
	right := sortedArray[middleIndex+1:]

	return &BSTNode{
		value: middle,
		left:  BuildBinarySearchTree(left),
		right: BuildBinarySearchTree(right),
	}
}

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

func (node *BSTNode) Height() int {
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
