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
func CommonAncestor2(nodeA, nodeB *BinaryTreeNode) *BinaryTreeNode {
	nodeAAncestor := nodeA

	for {
		if isAncestor(nodeAAncestor, nodeB) {
			break
		}

		nodeAAncestor = nodeAAncestor.parent
	}

	return nodeAAncestor
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
