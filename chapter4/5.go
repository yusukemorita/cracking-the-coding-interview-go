package main

func IsBinarySearchTree(node *BinaryTreeNode) bool {
	return isBinarySearchTree(node, nil, nil)
}

func isBinarySearchTree(node *BinaryTreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && *min > node.value {
		return false
	}

	if max != nil && *max < node.value {
		return false
	}

	if node.left != nil && node.left.value > node.value {
		return false
	}

	if node.right != nil && node.right.value < node.value {
		return false
	}

	return isBinarySearchTree(node.left, min, &node.value) && isBinarySearchTree(node.right, &node.value, max)

}
