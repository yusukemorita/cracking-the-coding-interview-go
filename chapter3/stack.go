package main

type Stack struct {
	head   *LinkedListNode[string]
	bottom *LinkedListNode[string]
	length int
}

func (s *Stack) push(value string) {
	s.head = &LinkedListNode[string]{
		value: value,
		next:  s.head,
	}

	s.length++

	if s.length == 1 {
		s.bottom = s.head
	}
}

func (s *Stack) pop() (string, bool) {
	if s.head == nil {
		return "", false
	}

	popped := s.head

	s.head = s.head.next
	s.length--

	if s.length == 0 {
		s.bottom = nil
	}

	return popped.value, true
}

func (s *Stack) peek() (string, bool) {
	if s.head == nil {
		return "", false
	}

	return s.head.value, true
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}

// for debugging
func (s *Stack) values() []string {
	if s.isEmpty() {
		return []string{}
	}

	return s.head.values()
}
