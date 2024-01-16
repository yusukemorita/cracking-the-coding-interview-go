package main

import "fmt"

// 1. iterate through the stack and move all items to the secondary stack, while extracting the max the
// 2. insert max into the original stack
// 3. move all items back into the original stack
// 4. repeat, but leave the sorted items in the original stack
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

func SortStack2(stack *GenericStack[int]) {
	// this stack will be sorted so that large numbers are on top
	sortedStack := GenericStack[int]{}

	for {
		fmt.Printf("stack: %v, sortedStack: %v\n", stack.values(), sortedStack.values())

		popped, ok := stack.pop()
		if !ok {
			// stack is empty
			break
		}

		// insert the popped value in the correctly sorted position
		// in sortedStack.
		// any items that are larger than the popped value will be moved back
		// to the stack.
		// these will be added back to the sortedStack in later iterations of the same
		// for loop.
		for {
			sortedStackTop, ok := sortedStack.peek()
			// if sortedStack is empty, push the popped value
			if !ok {
				sortedStack.push(popped)
				break
			}

			// when popped is bigger than sortedStackTop, push on top
			if popped >= sortedStackTop {
				sortedStack.push(popped)
				break
			}

			// when popped is smaller than sortedStackTop, we need to insert it
			// underneath sortedStackTop. Therefore, move sortedStackTop back to the
			// original stack.
			sortedStackTop, ok = sortedStack.pop()
			if !ok {
				panic("this shouldn't happen, as the value has already been peeked")
			}
			stack.push(sortedStackTop)
		}
	}

	// move items back into the original stack
	for {
		value, ok := sortedStack.pop()
		if !ok {
			break
		}

		stack.push(value)
	}
}
