package main

type LinkedListNode[T comparable] struct {
	value    T
	next     *LinkedListNode[T]
}

func (node LinkedListNode[T]) values() []T {
	if node.next == nil {
		return []T{node.value}
	}

	return append([]T{node.value}, node.next.values()...)
}
