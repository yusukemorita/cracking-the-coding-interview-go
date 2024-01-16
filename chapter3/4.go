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
	q.newOnTop.push(value)
}

func (q *MyQueue) transferToOldOnTop() {
	// only transfer from oldOnTop to newOnTop when oldOnTop
	// is empty, otherwise the order gets mixed up
	if !q.oldOnTop.isEmpty() {
		return
	}

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
	q.transferToOldOnTop()

	return q.oldOnTop.pop()
}

func (q *MyQueue) peek() (string, bool) {
	q.transferToOldOnTop()

	return q.oldOnTop.peek()
}
