package main

func DeleteMiddleNode(middleNode *LinkedListNode) {
	middleNode.value = middleNode.pointer.value
	middleNode.pointer = middleNode.pointer.pointer
}
