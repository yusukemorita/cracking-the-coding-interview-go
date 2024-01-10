package main

// solution 1, use a temporary buffer
// O(N)
func RemoveDuplicates1(node *LinkedListNode) {
	// use a map as a set
	encounteredValues := map[string]bool{}

	currentNode := node
	var previousNode *LinkedListNode = nil
	for {
		_, encountered := encounteredValues[currentNode.value]
		if encountered {
			// remove the current node from the list
			previousNode.pointer = currentNode.pointer
		} else {
			encounteredValues[currentNode.value] = true
		}

		if currentNode.pointer == nil {
			break
		}

		previousNode = currentNode
		currentNode = currentNode.pointer
	}
}

// solution 2, don't use a temporary buffer
// currentNode traversal takes N steps
// duplicateChecker traversal takes (N-1) + (N-2) + ... + 2 + 1 = N(N-1)/2
// -> O(N^2)
// therefore, solution 2 uses less additional space, but takes more time
func RemoveDuplicates2(head *LinkedListNode) {
	// list has 0 or 1 nodes, so no duplicates are possible
	if head == nil || head.pointer == nil {
		return
	}

	var currentNode *LinkedListNode = head
	// duplicateChecker will traverse ahead of the currentNode,
	// and check for duplicates
	var duplicateChecker *LinkedListNode = head.pointer
	var duplicateCheckerPrevious *LinkedListNode = head

	for {
		if duplicateChecker != nil && duplicateChecker.value == currentNode.value {
			// if duplicateChecker is pointing to a duplicate node,
			// delete the node
			duplicateCheckerPrevious.pointer = duplicateChecker.pointer
		}

		if duplicateChecker == nil {
			if currentNode.pointer == nil {
				// finished traversing
				break
			}

			// duplicate checker has reached the end of the list,
			// move currentNode forward
			currentNode = currentNode.pointer
			duplicateChecker = currentNode.pointer
			duplicateCheckerPrevious = currentNode
		} else {
			duplicateCheckerPrevious = duplicateChecker
			duplicateChecker = duplicateChecker.pointer
		}
	}
}
