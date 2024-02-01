package main

func BuildDepthLists(root *BSTNode) []*LinkedListNode[int] {
	var linkedLists []*LinkedListNode[int]
	nextLevel := []*BSTNode{root}

	for {
		if len(nextLevel) == 0 {
			break
		}

		var newNextLevel []*BSTNode
		var linkedList *LinkedListNode[int]

		for _, node := range nextLevel {
			nextItem := &LinkedListNode[int]{
				value:   node.value,
				pointer: linkedList,
			}
			linkedList = nextItem

			if node.left != nil {
				newNextLevel = append(newNextLevel, node.left)
			}
			if node.right != nil {
				newNextLevel = append(newNextLevel, node.right)
			}
		}

		linkedLists = append(linkedLists, linkedList)
		nextLevel = newNextLevel
	}

	return linkedLists
}

type LinkedListNode[T comparable] struct {
	value   T
	pointer *LinkedListNode[T]
}

func (node LinkedListNode[T]) values() []T {
	if node.pointer == nil {
		return []T{node.value}
	}

	return append([]T{node.value}, node.pointer.values()...)
}

func (node LinkedListNode[T]) length() int {
	if node.pointer == nil {
		return 1
	}

	return node.pointer.length() + 1
}
