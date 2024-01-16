package main

// queue implemented using two stacks
// FIFO
type MyQueue struct {
	// one stack is always empty
	newOnTop *Stack
	oldOnTop *Stack
}

func NewMyQueue() MyQueue {
	return MyQueue{
		newOnTop: &Stack{},
		oldOnTop: &Stack{},
	}
}

// O(1)
func (q *MyQueue) push(value string) {
	if q.newOnTop.isEmpty() && !q.oldOnTop.isEmpty() {
		q.transferToNewOnTop()
	}

	q.newOnTop.push(value)
}

func (q *MyQueue) transferToNewOnTop() {
	// move everything to newOnTop
	for {
		popped, ok := q.oldOnTop.pop()
		if !ok {
			break
		}

		q.newOnTop.push(popped)
	}
}

func (q *MyQueue) transferToOldOnTop() {
	// move everything to oldOnTop
	for {
		popped, ok := q.newOnTop.pop()
		if !ok {
			break
		}

		q.oldOnTop.push(popped)
	}
}

// O(2N) = O(N), as all items in newOnTop have to be
// moved to oldOnTop and then back
func (q *MyQueue) pop() (string, bool) {
	if !q.newOnTop.isEmpty() && q.oldOnTop.isEmpty() {
		q.transferToOldOnTop()
	}

	return q.oldOnTop.pop()
}

func (q *MyQueue) peek() (string, bool) {
	if !q.newOnTop.isEmpty() && q.oldOnTop.isEmpty() {
		q.transferToOldOnTop()
	}

	return q.oldOnTop.peek()
}
