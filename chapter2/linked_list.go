package main

type GenericLinkedListNode[T comparable] struct {
	value   T
	pointer *GenericLinkedListNode[T]
}

func (node GenericLinkedListNode[T]) values() []T {
	if node.pointer == nil {
		return []T{node.value}
	}

	return append([]T{node.value}, node.pointer.values()...)
}

type LinkedListNode struct {
	value   string
	pointer *LinkedListNode
}

func (node LinkedListNode) length() int {
	if node.pointer == nil {
		return 1
	}

	return node.pointer.length() + 1
}

func (node LinkedListNode) values() []string {
	if node.pointer == nil {
		return []string{node.value}
	}

	return append([]string{node.value}, node.pointer.values()...)
}
