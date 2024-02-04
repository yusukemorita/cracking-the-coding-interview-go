package main

import "fmt"

func BSTSequence(root *BinaryTreeNode) {
	bstSequence(root, []int{}, []*BinaryTreeNode{})
}

func bstSequence(root *BinaryTreeNode, result []int, nextCandidates []*BinaryTreeNode) {
	// debug
	var nextValues []int
	for _, c := range nextCandidates {
		nextValues = append(nextValues, c.value)
	}
	fmt.Printf("root: %d, result: %v, nextCandidates: %v\n", root.value, result, nextValues)
	// debug end

	if root != nil {
		if root.left != nil {
			// fmt.Printf("adding %d to next candidates\n", root.left.value)
			nextCandidates = append(nextCandidates, root.left)
		}
		if root.right != nil {
			// fmt.Printf("adding %d to next candidates\n", root.right.value)
			nextCandidates = append(nextCandidates, root.right)
		}
	}

	if root != nil {
		result = append(result, root.value)
	}

	if len(nextCandidates) == 0 {
		fmt.Println(result)
		return
	}

	for _, current := range nextCandidates {
		nextCandidatesForCurrent := removeFromSlice(nextCandidates, current)
		bstSequence(current, result, nextCandidatesForCurrent)
	}
}

func removeFromSlice[T comparable](slice []T, toRemove T) []T {
	var newSlice []T

	for _, item := range slice {
		if item != toRemove {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}
