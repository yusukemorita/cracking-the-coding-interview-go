package main

type Node = GenericLinkedListNode[int]

// use O(N) additional space
func Partition(head *Node, border int) {
	// a: smaller than border
	// b: greater than or equal to border
	//
	// e.g.
	// a linked list with items (a, b, a, b, a)
	// must be sorted into (a, a, a, b, b)

	// keep all b nodes in a slice, so they can be swapped out if necessary
	// use like a FIFO queue
	aboveBorderNodes := []*Node{}

	current := head
	for {
		if current == nil {
			break
		}

		// if the current node is an "a" and there are "b" nodes encountered,
		// then swap the current "a" node with the first encountered "b" node
		if current.value < border && len(aboveBorderNodes) > 0 {
			bValue := aboveBorderNodes[0].value
			aboveBorderNodes[0].value = current.value
			current.value = bValue

			// remove the first node from aboveBorderNodes, as it now contains a
			// value that is below the border
			aboveBorderNodes = aboveBorderNodes[1:]

		}

		if current.value >= border {
			aboveBorderNodes = append(aboveBorderNodes, current)
		}

		current = current.pointer
	}
}
