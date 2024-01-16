package main

// queue implemented using two stacks
// FIFO
type MyQueue struct {
	// one stack is always empty
	mainStack *Stack
	tempStack *Stack
}

func NewMyQueue() MyQueue {
	return MyQueue{
		mainStack: &Stack{},
		tempStack: &Stack{},
	}
}

// O(1)
func (q *MyQueue) push(value string) {
	q.mainStack.push(value)
}

// O(2N) = O(N), as all items in main stack have to be
// moved to temp stack and then back
func (q *MyQueue) pop() (string, bool) {
	// first, transfer everything to the tempStack
	for {
		popped, ok := q.mainStack.pop()
		if !ok {
			break
		}

		q.tempStack.push(popped)
	}

	poppedFromTemp, ok := q.tempStack.pop()
	if !ok {
		// main stack was empty to begin with
		return "", false
	}

	// return everything back to the main stack
	for {
		popped, ok := q.tempStack.pop()
		if !ok {
			break
		}

		q.mainStack.push(popped)
	}

	return poppedFromTemp, true
}

func (q *MyQueue) peek() (string, bool) {
	// first, transfer everything to the tempStack
	for {
		popped, ok := q.mainStack.pop()
		if !ok {
			break
		}

		q.tempStack.push(popped)
	}

	peekedFromTemp, ok := q.tempStack.peek()
	if !ok {
		// main stack was empty to begin with
		return "", false
	}

	// return everything back to the main stack
	for {
		popped, ok := q.tempStack.pop()
		if !ok {
			break
		}

		q.mainStack.push(popped)
	}

	return peekedFromTemp, true
}