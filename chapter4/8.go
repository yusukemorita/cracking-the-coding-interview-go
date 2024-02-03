package main

// with links to parents
func CommonAncestor(nodeA, nodeB *BinaryTreeNode) *BinaryTreeNode {
	nodeAAncestors := make(map[*BinaryTreeNode]bool)
	nodeBAncestors := make(map[*BinaryTreeNode]bool)

	nextNodeAAncestor := nodeA
	nextNodeBAncestor := nodeB

	for {
		if nextNodeAAncestor != nil {
			if _, ok := nodeBAncestors[nextNodeAAncestor]; ok {
				return nextNodeAAncestor
			}

			nodeAAncestors[nextNodeAAncestor] = true
			nextNodeAAncestor = nextNodeAAncestor.parent
		}

		if nextNodeBAncestor != nil {
			if _, ok := nodeAAncestors[nextNodeBAncestor]; ok {
				return nextNodeBAncestor
			}

			nodeBAncestors[nextNodeBAncestor] = true
			nextNodeBAncestor = nextNodeBAncestor.parent
		}
	}
}

// with links to parents
// avoid storing additional nodes in a data structure
// time: O((logN)^2)
func CommonAncestor2(root, nodeA, nodeB *BinaryTreeNode) *BinaryTreeNode {
	ancestor := root

	for {
		if isAncestor(ancestor.left, nodeA) && isAncestor(ancestor.left, nodeB) {
			// left is an ancestor of both nodes
			ancestor = ancestor.left
			continue
		}

		if isAncestor(ancestor.right, nodeA) && isAncestor(ancestor.right, nodeB) {
			// right is an ancestor of both nodes
			ancestor = ancestor.right
			continue
		}

		// neither left or right are both ancestors, which means that the current `ancestor`
		// is the first common ancestor
		break
	}

	return ancestor
}

// O(logN)
func isAncestor(maybeAncestor, node *BinaryTreeNode) bool {
	if node == nil {
		return false
	}

	if node == maybeAncestor {
		return true
	}

	return isAncestor(maybeAncestor, node.parent)
}
