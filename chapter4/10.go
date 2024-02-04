package main

func IsSubtree(tree, maybeSubtree *BinaryTreeNode) bool {
	// treat `nil` as not a subtree of anything
	if tree == nil || maybeSubtree == nil {
		return false
	}

	return IsIdentical(tree, maybeSubtree) ||
		IsSubtree(tree.left, maybeSubtree) ||
		IsSubtree(tree.right, maybeSubtree)
}

func IsIdentical(tree1, tree2 *BinaryTreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil && tree2 != nil || tree1 != nil && tree2 == nil {
		return false
	}

	return tree1.value == tree2.value &&
		IsIdentical(tree1.left, tree2.left) &&
		IsIdentical(tree1.right, tree2.right)
}
