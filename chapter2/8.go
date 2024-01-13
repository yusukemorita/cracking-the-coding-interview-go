package main

// store all encountered memory addresses, and return the first
// memory address that we run into that is already encountered
// uses O(N)
func DetectLoop(head *LinkedListNode) *LinkedListNode {
	encounteredNodes := map[*LinkedListNode]bool{}

	current := head
	for {
		if current == nil {
			break
		}

		_, ok := encounteredNodes[current]
		if ok {
			return current
		}

		encounteredNodes[current] = true

		current = current.pointer
	}

	return nil
}

// different implementation: use runners
// uses O(1) space
func DetectLoop2(head *LinkedListNode) *LinkedListNode {
	// move fastRunner twice as fast as slowRunner
	// if there is a loop, they will eventually collide
	// if not, fastRunner will reach the end of the list
	slowRunner := head
	fastRunner := head
	var collisionSpot *LinkedListNode

	for {
		// move slowRunner 1, fastRunner 2
		slowRunner = slowRunner.pointer
		if fastRunner.pointer == nil {
			// reached the end of the list, no loop
			return nil
		}
		fastRunner = fastRunner.pointer.pointer

		if fastRunner == nil {
			// reached the end of the list, no loop
			return nil
		}

		if slowRunner == fastRunner {
			// found a loop
			collisionSpot = slowRunner
			break
		}

	}

	// WHEN SLOW RUNNER ENTERS THE LOOP
	// If K = the distance from the head to the start of the loop till
	// distance traveled by slowRunner: K
	// distance traveled by fastRunner: 2K
	//
	// This means fastRunner is K positions into the loop, and the distance between
	// fastRunner and slowRunner is (loopSize - K) steps.
	// Since fastRunner moves twice as fast as slowRunner, it moves closer
	// to slowRunner by 1 position / unit of time.
	// Therefore, slowRunner and fastRunner will collide after an additional
	// (loopSize - K) steps.
	// -> the collision point will be (loopSize - K) steps into the loop,
	//    or K steps from the end/start of the loop
	// since K is also the number of steps from the head to the start of the loop,
	// we can find the position of the start of the loop by
	// running two pointers from the head and the collision point,
	// and returning the position that they meet each other.
	headRunner := head
	collisionRunner := collisionSpot

	for {
		if headRunner == collisionRunner {
			// this is the start of the loop
			return headRunner
		}

		headRunner = headRunner.pointer
		collisionRunner = collisionRunner.pointer
	}
}
