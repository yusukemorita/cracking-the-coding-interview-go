package main

import (
	"testing"
)

func TestDetectLoop(t *testing.T) {
	t.Run("returns start of loop when there is a loop", func(t *testing.T) {
		startOfLoop := &LinkedListNode{
			value: "loop1",
			pointer: &LinkedListNode{
				value: "loop2",
			},
		}

		startOfLoop.pointer.pointer = startOfLoop

		linkedList := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value:   "hello",
				pointer: startOfLoop,
			},
		}

		got := DetectLoop(&linkedList)

		if got != startOfLoop {
			t.Errorf("expected %v, got %v", startOfLoop, got)
		}
	})

	t.Run("returns nil when there is no loop", func(t *testing.T) {
		linkedList := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "hello",
				pointer: &LinkedListNode{
					value: "bonjour",
				},
			},
		}

		got := DetectLoop(&linkedList)

		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
}

func TestDetectLoop2(t *testing.T) {
	t.Run("returns start of loop when there is a loop", func(t *testing.T) {
		startOfLoop := &LinkedListNode{
			value: "loop1",
			pointer: &LinkedListNode{
				value: "loop2",
			},
		}

		startOfLoop.pointer.pointer = startOfLoop

		linkedList := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value:   "hello",
				pointer: startOfLoop,
			},
		}

		got := DetectLoop2(&linkedList)

		if got != startOfLoop {
			t.Errorf("expected %v, got %v", startOfLoop, got)
		}
	})

	t.Run("returns nil when there is no loop", func(t *testing.T) {
		linkedList := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "hello",
				pointer: &LinkedListNode{
					value: "bonjour",
				},
			},
		}

		got := DetectLoop2(&linkedList)

		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
}
