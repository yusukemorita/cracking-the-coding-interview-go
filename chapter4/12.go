package main

// each node will be visited "depth" amount of times,
// so time complexity is O(NlogN).
// Since no additional data structures are used and recursion is used, space complexity is O(logN)
func CountPathsWithSum(node *BinaryTreeNode, sum int) int {
	result := 0

	// when path ends with current node
	if node.value == sum {
		result++
	}

	if node.left != nil {
		// when path does not include current node, and starts with left
		result += CountPathsWithSum(node.left, sum)

		if node.value != 0 {
			// when path includes current node
			result += CountPathsWithSum(node.left, sum-node.value)
		}
	}

	if node.right != nil {
		// when path does not include current node, and starts with right
		result += CountPathsWithSum(node.right, sum)

		if node.value != 0 {
			// when path includes current node
			result += CountPathsWithSum(node.right, sum-node.value)
		}
	}

	return result
}
