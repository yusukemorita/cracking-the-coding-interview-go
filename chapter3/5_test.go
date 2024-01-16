package main

import (
	"testing"
)

func TestSortStack(t *testing.T) {
	t.Run("does not raise error when stack is empty", func(t *testing.T) {
		stack := GenericStack[int]{}
		SortStack(&stack)
	})

	t.Run("does nothing when stack has 1 item", func(t *testing.T) {
		stack := GenericStack[int]{}
		stack.push(1)
		SortStack(&stack)

		popped, ok := stack.pop()
		if !ok || popped != 1 {
			t.Errorf("expected ok: true, popped: 1, got ok: %v, value: %d", ok, popped)
		}

		_, ok = stack.pop()
		if ok {
			t.Error("expected ok: false")
		}
	})

	t.Run("sorts stack when stack has more than 1 item", func(t *testing.T) {
		stack := GenericStack[int]{}
		stack.push(2)
		stack.push(1)
		stack.push(5)
		stack.push(3)
		stack.push(4)

		SortStack(&stack)

		for _, expected := range []int{1, 2, 3, 4, 5} {
			value, ok := stack.pop()
			if !ok || value != expected {
				t.Errorf("expected ok: true, value: %d, got ok: %v, value: %d", expected, ok, value)
			}
		}

		_, ok := stack.pop()
		if ok {
			t.Error("expected ok: false")
		}
	})
}
