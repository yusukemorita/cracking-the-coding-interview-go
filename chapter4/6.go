package main

// find next (in-order successor) node of a binary search tree
// for simplicity, assume all node values are unique.
func NextNode(node *BinaryTreeNode) *BinaryTreeNode {
	if node.right == nil {
		// no parent and no right
		//
		//   10 <- this case
		//  5
		// 3 6
		if node.parent == nil {
			return nil
		}

		// no right, but has a parent that is bigger than itself
		//
		//       10
		//   -> 5  20
		//  -> 2
		//
		if node.parent.value > node.value {
			return node.parent
		}

		// if parent is smaller than node, then keep traveling up
		// until a larger node is found. If no larger node is found,
		// then the current node is the largest node, so return nil.
		//
		//     10
		//   5    15
		//  3 7  13 20 <-
		//   6 8 <-
		return findAncestorLargerThanValue(node, node.value)
	}

		//     10 <-
		//-> 5    15 <-
		//  3 7  13 20
		//   6 8
	return findMostLeftDescendant(node.right)
}

func findAncestorLargerThanValue(node *BinaryTreeNode, value int) *BinaryTreeNode {
	if node.parent == nil {
		return nil
	}

	if node.parent.value > value {
		return node.parent
	}

	return findAncestorLargerThanValue(node.parent, value)
}

func findMostLeftDescendant(node *BinaryTreeNode) *BinaryTreeNode {
	if node.left == nil {
		return node
	}

	return findMostLeftDescendant(node.left)
}
