package main

import "fmt"

// assume numbers in stack are greater than 0
func SortStack(stack *GenericStack[int]) {
	var max *int
	tempStack := GenericStack[int]{}
	count := 0

	for {
		popped, ok := stack.pop()
		if !ok {
			break
		}

		count++

		if max == nil {
			max = &popped
		} else if popped > *max {
			// ignore the default value of max
			tempStack.push(*max)
			max = &popped
		} else {
			tempStack.push(popped)
		}

		fmt.Printf("stack: %v, tempStack: %v, max: %d\n", stack.values(), tempStack.values(), *max)
	}

	// no need to do anything if length is 0
	if count == 0 {
		return
	}

	stack.push(*max)
	max = nil
	fmt.Printf("stack: %v, tempStack: %v\n", stack.values(), tempStack.values())

	// move everything back to the stack
	for {
		popped, ok := tempStack.pop()
		if !ok {
			break
		}

		stack.push(popped)
	}
	fmt.Printf("stack: %v, tempStack: %v\n", stack.values(), tempStack.values())

	sortedCount := 1

	// repeat the process until everything is sorted
	for {
		if sortedCount == count {
			break
		}

		// move everything (except the sorted items) to the tempstack and extract the max
		for i := 1; i <= count-sortedCount; i++ {
			popped, ok := stack.pop()
			if !ok {
				panic("should not happen")
			}

			if max == nil {
				max = &popped
			} else if popped > *max {
				// ignore the default value of max
				tempStack.push(*max)
				max = &popped
			} else {
				tempStack.push(popped)
			}

			fmt.Printf("stack: %v, tempStack: %v, max: %d\n", stack.values(), tempStack.values(), *max)
		}

		// move the max back into the stack
		stack.push(*max)
		max = nil
		fmt.Printf("stack: %v, tempStack: %v\n", stack.values(), tempStack.values())

		// move everything else back into the stack
		for {
			popped, ok := tempStack.pop()
			if !ok {
				break
			}

			stack.push(popped)
		}
		fmt.Printf("stack: %v, tempStack: %v\n", stack.values(), tempStack.values())

		// increment the sortedCount
		sortedCount++
	}
}
