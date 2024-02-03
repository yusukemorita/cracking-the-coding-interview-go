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

// without links to parents
// avoid storing additional nodes in a data structure
// time: O(n)
func CommonAncestor4(root, nodeA, nodeB *BinaryTreeNode) *BinaryTreeNode {
	ancestor := root
	if root == nodeA || root == nodeB {
		return root
	}

	for {
		// Assuming tree is balanced, visits all nodes on the left side twice,
		// so visit count = 2 * n/2 = n
		// As the loop progresses, each time the number of nodes visited is cut in
		// approximately half. So, the number of nodes visited is
		//      n + n/2 + n/4 + ... + 2 + 1 = X
		// 2n + n + n/2 + n/4 + ... + 2     = 2X
		// -> X = 2n + 1
		// therefore, the time complexity is O(2n-1) = O(n)
		aIsOnLeft := covers(ancestor.left, nodeA)
		bIsOnLeft := covers(ancestor.left, nodeB)

		if aIsOnLeft && bIsOnLeft {
			ancestor = root.left
			continue
		}

		if !aIsOnLeft && !bIsOnLeft {
			ancestor = root.right
			continue
		}

		// neither left nor right are common ancestors, so the current `ancestor`
		// is the first common ancestor
		break
	}

	return ancestor
}


func covers(maybeAncestor, node *BinaryTreeNode) bool {
	if maybeAncestor == nil {
		return false
	}

	if maybeAncestor == node {
		return true
	}

	return covers(maybeAncestor.left, node) || covers(maybeAncestor.right, node)
}

func CommonAncestor5(root, nodeA, nodeB *BinaryTreeNode) (commonAncestor *BinaryTreeNode, hasA, hasB bool) {
	if root == nil {
		return nil, false, false
	}

	commonLeft, leftHasA, leftHasB := CommonAncestor5(root.left, nodeA, nodeB)
	if commonLeft != nil {
		return commonLeft, true, true
	}

	commonRight, rightHasA, rightHasB := CommonAncestor5(root.right, nodeA, nodeB)
	if commonRight != nil {
		return commonRight, true, true
	}

	hasA = leftHasA || rightHasA || root == nodeA
	hasB = leftHasB || rightHasB || root == nodeB

	if hasA && hasB {
		return root, true, true
	}

	return nil, hasA, hasB
}
