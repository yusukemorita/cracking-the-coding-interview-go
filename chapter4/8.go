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
// time: O((logN)^2) for a balanced tree
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

// with links to parents
// avoid storing additional nodes in a data structure
// improve performance over CommonAncestor2
// time: O(N), since every node is searched only once
func CommonAncestor3(nodeA, nodeB *BinaryTreeNode) *BinaryTreeNode {
	current := nodeA
	if search(current, nodeB) {
		return current
	}

	for {
		if current.parent == nil {
			// current is root
			return current
		}

		parent := current.parent
		if parent == nodeB {
			return parent
		}

		var sibling *BinaryTreeNode
		if parent.left == current {
			sibling = parent.right
		} else {
			sibling = parent.left
		}
		if search(sibling, nodeB) {
			return parent
		}

		current = current.parent
	}
}

// return true if the target is found in the node subtree
func search(node, target *BinaryTreeNode) bool {
	if node == nil {
		return false
	}

	if node == target {
		return true
	}

	return search(node.left, target) || search(node.right, target)
}
