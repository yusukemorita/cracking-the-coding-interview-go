package main

import (
	"errors"
)

func KthToLast1(head *LinkedListNode, k int) (*LinkedListNode, error) {
	if k <= 0 {
		return nil, errors.New("k must be greater than 0")
	}

	// we want to use two pointers, (k - 1) positions apart, and move them together until
	// the pointer ahead reaches the end.

	// first, move the ahead (k - 1) times
	ahead := head
	aheadIndex := 0
	for {
		if aheadIndex == k-1 {
			break
		}

		if ahead == nil {
			return nil, errors.New("k must be less than or equal to the length of the linked list")
		}

		ahead = ahead.pointer
		aheadIndex++
	}

	// now, behind is at index 0, ahead is at index k - 1, and we move
	// ahead and behind together.
	// if the length of the linked list is N, when ahead reaches the last node (index N - 1),
	// behind will be at index (N - k + 1), which is the kth index from the end.
	// e.g. If a list has 4 items (N = 4) and k = 2,
	//      - ahead will be at index 3, which is the end
	//      - behind will be at index (3 - 2 + 1) = 2, which is 2nd from the end
	behind := head

	for {
		if ahead == nil {
			return nil, errors.New("k must be less than or equal to the length of the linked list")
		}

		if ahead.pointer == nil {
			break
		}

		ahead = ahead.pointer
		behind = behind.pointer
	}

	return behind, nil
}

// recursive solution

type Counter struct {
	value int
}

func (c *Counter) increment() {
	c.value++
}

func KthToLast2(head *LinkedListNode, k int, counter *Counter) *LinkedListNode {
	if head == nil {
		return nil
	}

	node := KthToLast2(head.pointer, k, counter)
	counter.increment()

	if counter.value == k {
		return head
	}

	return node
}
