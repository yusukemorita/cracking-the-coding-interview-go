package main

// if there is an intersection, then all nodes after the intersection
// are also equal.
// therefore, the last node should also be equal. both lists should mean that all nodes
func FindIntersection(l1, l2 *LinkedListNode) *LinkedListNode {
	// find the lengths of each list
	// O(A+B)
	// A: length of L1
	// B: length of L2
	l1Length := l1.length()
	l2Length := l2.length()

	currentL1 := l1
	currentL2 := l2

	// If there is an intersection, the last N (>= 1) nodes of each list are
	// identical. Therefore, if we start traversing the last K nodes of each
	// list together, we will eventually find identical nodes.
	// To do this, we must first adjust for the difference in length.
	if l1Length >= l2Length {
		for i := 0; i < l1Length-l2Length; i++ {
			currentL1 = currentL1.pointer
		}
	} else {
		for i := 0; i < l2Length-l1Length; i++ {
			currentL2 = currentL2.pointer
		}
	}

	for {
		if currentL1 == nil || currentL2 == nil {
			break
		}

		if currentL1 == currentL2 {
			// this is the first identical node = intersection
			return currentL1
		}

		currentL1 = currentL1.pointer
		currentL2 = currentL2.pointer
	}

	// no intersection found
	return nil
}
