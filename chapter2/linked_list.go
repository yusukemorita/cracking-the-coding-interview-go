package main

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
