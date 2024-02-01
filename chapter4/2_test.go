package main

import (
	"fmt"
	"testing"
)

// difficult to test, so just print the BST
func TestBuildBinarySearchTree(t *testing.T) {
	t.Run("when there are lots of values", func(t *testing.T) {
		sortedArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		bst := BuildBinarySearchTree(sortedArray)
		printBST(bst)
	})

	t.Run("when there are no values", func(t *testing.T) {
		sortedArray := []int{}
		BuildBinarySearchTree(sortedArray)
	})

	t.Run("when there is one value", func(t *testing.T) {
		sortedArray := []int{1}
		bst := BuildBinarySearchTree(sortedArray)
		printBST(bst)
	})
}

func printBST(node *BSTNode) {
	nextLevel := []*BSTNode{node}

	for {
		if len(nextLevel) == 0 {
			break
		}

		var newNextLevel []*BSTNode

		for _, n := range nextLevel {
			fmt.Printf(" %d ", n.value)
			if n.left != nil {
				newNextLevel = append(newNextLevel, n.left)
			}
			if n.right != nil {
				newNextLevel = append(newNextLevel, n.right)
			}
		}

		fmt.Print("\n")
		nextLevel = newNextLevel
	}
}
