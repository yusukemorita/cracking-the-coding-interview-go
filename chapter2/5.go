package main

import "fmt"

// number A: 123 (3 -> 2 -> 1)
// number B: 45 (5 -> 4)
// output: 168 (8 -> 6 -> 1)
func AddNumbersReversed(numberA *Node, numberB *Node) *Node {
	// assume numberA and numberB are not nil
	var head *Node
	var last *Node

	currentA := numberA
	currentB := numberB
	carryOver := 0

	for {
		if currentA == nil && currentB == nil {
			break
		}

		value := carryOver

		if currentA != nil {
			value += currentA.value
			currentA = currentA.pointer
		}

		if currentB != nil {
			value += currentB.value
			currentB = currentB.pointer
		}

		fmt.Printf("value before carryOver: %d\n", value)

		if value >= 10 {
			value = value - 10
			carryOver = 1
		} else {
			carryOver = 0
		}

		fmt.Printf("value: %d, carryOver: %d\n", value, carryOver)

		if last == nil {
			last = &GenericLinkedListNode[int]{
				value: value,
			}
			head = last

		} else {
			last.pointer = &GenericLinkedListNode[int]{
				value: value,
			}
			last = last.pointer
		}
	}

	return head
}

// number A: 123 (1 -> 2 -> 3)
// number B: 45 (4 -> 5)
// output: 168 (1 -> 6 -> 8)
func AddNumbers(numberA, numberB *Node) *Node {
	calculatedA := 0
	currentA := numberA
	for {
		if currentA == nil {
			break
		}

		calculatedA = 10*calculatedA + currentA.value
		currentA = currentA.pointer
	}

	calculatedB := 0
	currentB := numberB
	for {
		if currentB == nil {
			break
		}

		calculatedB = 10*calculatedB + currentB.value
		currentB = currentB.pointer
	}

	sum := calculatedA + calculatedB

	var linkedList *Node

	for {
		if sum == 0 {
			break
		}

		linkedList = &Node{
			value:   sum % 10,
			pointer: linkedList,
		}

		sum = sum / 10
	}

	return linkedList
}
