package main

import (
	"errors"
	"fmt"
)

func KthToLast1(head *LinkedListNode, k int) (*LinkedListNode, error) {
	if k <= 0 {
		return nil, errors.New("k must be greater than 0")
	}

	// first, count the length
	length := 0
	current := head
	for {
		if current == nil {
			break
		} else {
			length++
			current = current.pointer
		}
	}

	if k > length {
		return nil, errors.New("k must be smaller than the length of the linked list")
	}

	// assume item at index (length - k) exists
	currentIndex := 0
	targetIndex := length - k
	current = head

	for {
		if currentIndex == targetIndex {
			break
		}

		current = current.pointer
		currentIndex++
	}

	return current, nil
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
