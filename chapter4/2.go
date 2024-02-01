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
