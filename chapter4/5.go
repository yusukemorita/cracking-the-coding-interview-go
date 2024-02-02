package main

// a binary tree is a binary search tree when, for any node,
// all nodes in the left subtree are smaller than or equal to the node value,
// and all nodes in the right subtree are larger than the node value
// i.e. (any left subtree node) <= node < (any right subtree node)
func IsBinarySearchTree(node *BinaryTreeNode) bool {
	return isBinarySearchTree(node, nil, nil)
}

// min is inclusive, max is exclusive
func isBinarySearchTree(node *BinaryTreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.value <= *min {
		return false
	}

	if max != nil && *max < node.value {
		return false
	}

	return isBinarySearchTree(node.left, min, &node.value) && isBinarySearchTree(node.right, &node.value, max)
}
