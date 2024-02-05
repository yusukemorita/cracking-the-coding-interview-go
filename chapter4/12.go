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

// O(N) time complexity, as each node is visited once
// O(N) space complexity, due to the pathCount map
func CountPathsWithSum2(node *BinaryTreeNode, targetSum int) int {
	pathCount := make(map[int]int)
	// if a path with a sum of `targetSum` starts at root, then there needs to be
	// a value for "sum = zero" in the map to count this
	pathCount[0] = 1
	return countPathsWithSum(node, targetSum, 0, pathCount)
}

func countPathsWithSum(node *BinaryTreeNode, targetSum int, runningSum int, pathCount map[int]int) int {
	if node == nil {
		return 0
	}

	runningSum += node.value
	pathCount[runningSum]++

	result := pathCount[runningSum-targetSum]

	result += countPathsWithSum(node.left, targetSum, runningSum, pathCount)
	result += countPathsWithSum(node.right, targetSum, runningSum, pathCount)

	pathCount[runningSum]--

	return result
}
