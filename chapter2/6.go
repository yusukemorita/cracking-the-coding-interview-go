package main

// a palindrome is a word that is the same read
// forwards and backward

// compare each node value, don't concatenate together
func IsPalindrome(head *LinkedListNode) bool {
	if head == nil {
		// return false when linked list is nil. this might change depending on the
		// requirements
		return false
	}

	// find the length, and also disconnect the last node
	length := 0
	current := head
	var previous *LinkedListNode // ends up with second to last node
	var last *LinkedListNode

	for {
		length++

		if current.pointer == nil {
			last = current
			break
		}

		previous = current
		current = current.pointer
	}

	if length == 1 {
		return true
	}

	// disconnect the last node.
	// If we were given a linked list like a -> b -> c -> d -> e, then
	// in the next recursive call we should compare b -> c -> d only
	previous.pointer = nil

	return head.value == last.value && IsPalindrome(head.pointer)
}

// the first solution computes the length N/2 times
// IsPalindrome can be implemented more efficiently using Reverse()
func IsPalindrome2(head *LinkedListNode) bool {
	if head == nil {
		// return false when linked list is nil. this might change depending on the
		// requirements
		return false
	}

	// clone and reverse the linked list
	// clone is necessary because Reverse() modifies the linked list
	headClone := *head
	reversed := ReverseWithPointers(&headClone)

	// only need to compare the first half of each list to know if it is
	// a palindrome
	// however, the length is not known here, so compare all items
	// we could reimplement `Reverse()` to return the length too
	currentOriginal := head
	currentReversed := reversed
	for {
		// since original and reversed have the same length, they should become
		// nil at the same time
		if currentOriginal == nil && currentReversed == nil {
			break
		}

		if currentOriginal.value != currentReversed.value {
			return false
		}

		currentOriginal = currentOriginal.pointer
		currentReversed = currentReversed.pointer
	}

	return true
}

func Reverse(head *LinkedListNode) (reversedHead, reversedTail *LinkedListNode) {
	if head.pointer == nil {
		// only 1 node in list, no need to do anything
		return head, head
	}

	rh, rt := Reverse(head.pointer)
	rt.pointer = head
	head.pointer = nil
	return rh, head
}

// not recursive implementation
func ReverseWithPointers(head *LinkedListNode) *LinkedListNode {
	var previous *LinkedListNode
	current := head

	for {
		if current == nil {
			break
		}

		// save the next node temporarily, as current.pointer
		// is overwritten in the subsequent code
		next := current.pointer

		// reverse direction of current pointer
		current.pointer = previous

		// move pointers one forward
		previous = current
		current = next
	}

	// previous ends up with the new head
	return previous
}
