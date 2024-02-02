package main

func IsBalanced1(node *BinaryTreeNode) bool {
	if node == nil {
		return true
	}

	leftHeight := 0
	if node.left != nil {
		leftHeight = node.left.Height()
	}

	rightHeight := 0
	if node.right != nil {
		rightHeight = node.right.Height()
	}

	diff := leftHeight - rightHeight
	if diff >= 2 || diff <= -2 {
		return false
	}

	return IsBalanced1(node.left) && IsBalanced1(node.right)
}

// implement without using the Height() method
func IsBalanced(node *BinaryTreeNode) (result bool, height int) {
	if node == nil {
		return true, 0
	}

	leftResult, leftHeight := IsBalanced(node.left)
	rightResult, rightHeight := IsBalanced(node.right)

	height = max(leftHeight, rightHeight) + 1

	diff := leftHeight - rightHeight
	if diff >= 2 || diff <= -2 {
		result = false
	} else {
		result = leftResult && rightResult
	}

	return
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
