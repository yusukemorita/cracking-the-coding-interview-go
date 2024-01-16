package main

type SetOfStacks struct {
	stacks        []*Stack
	stackCapacity uint
}

func (s *SetOfStacks) currentStack() *Stack {
	return s.stacks[len(s.stacks)-1]
}

func (s *SetOfStacks) push(value string) {
	if len(s.stacks) == 0 {
		s.stacks = append(s.stacks, &Stack{})
	}

	if s.currentStack().length == int(s.stackCapacity) {
		s.stacks = append(s.stacks, &Stack{})
	}

	s.currentStack().push(value)
}

func (s *SetOfStacks) pop() (string, bool) {
	if len(s.stacks) == 0 {
		return "", false
	}

	value, ok := s.currentStack().pop()

	if s.currentStack().length == 0 {
		// remove last stack because it is now empty
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return value, ok
}

func (s *SetOfStacks) popSubStack(index uint) (string, bool) {
	subStack := s.stacks[index]

	value, ok := subStack.pop()

	s.removeEmptySubStacks()

	return value, ok
}

// shift items in substacks so that all substacks except the
// last are using maximum capacity
// func (s *SetOfStacks) reorganize() {
// 	for i, subStack := range s.stacks {
// 		// ignore last stack, as this does not need to be full
// 		if i == len(s.stacks)-1 {
// 			continue
// 		}

// 		if subStack.length < int(s.stackCapacity) {
// 			nextStack := s.stacks[i+1]

// 		}

// 	}
// }

func (s *SetOfStacks) removeEmptySubStacks() {
	// traverse s.stacks backwards, to prevent the index from shifting
	// when a stack is removed
	for i := len(s.stacks) - 1; i >= 0; i-- {
		subStack := s.stacks[i]

		if subStack.length == 0 {
			// remove the empty substack
			s.stacks = append(s.stacks[0:i], s.stacks[i+1:]...)
		}
	}
}
