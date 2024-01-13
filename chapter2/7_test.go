package main

import (
	"testing"
)

func TestFindIntersection(t *testing.T) {
	t.Run("returns the intersection when in the middle", func(t *testing.T) {
		intersection := &LinkedListNode{
			value: "intersection",
			pointer: &LinkedListNode{
				value: "intersection2",
			},
		}

		l1 := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "hello",
				pointer: &LinkedListNode{
					value:   "hi",
					pointer: intersection,
				},
			},
		}

		l2 := LinkedListNode{
			value:   "hi",
			pointer: intersection,
		}

		got := FindIntersection(&l1, &l2)

		if got != intersection {
			t.Errorf("got %v", got)
		}
	})

	t.Run("returns the intersection when last", func(t *testing.T) {
		intersection := &LinkedListNode{
			value: "intersection",
		}

		l1 := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "hello",
				pointer: &LinkedListNode{
					value:   "hi",
					pointer: intersection,
				},
			},
		}

		l2 := LinkedListNode{
			value:   "hi",
			pointer: intersection,
		}

		got := FindIntersection(&l1, &l2)

		if got != intersection {
			t.Errorf("got %v", got)
		}
	})

	t.Run("returns nil when intersection is not present", func(t *testing.T) {
		l1 := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "hello",
				pointer: &LinkedListNode{
					value: "hi",
				},
			},
		}

		l2 := LinkedListNode{
			value: "hi",
		}

		got := FindIntersection(&l1, &l2)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("returns nil when values are the same but reference is different", func(t *testing.T) {
		l1 := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "hello",
				pointer: &LinkedListNode{
					value: "intersection",
				},
			},
		}

		l2 := LinkedListNode{
			value: "hi",
			pointer: &LinkedListNode{
				value: "intersection",
			},
		}

		got := FindIntersection(&l1, &l2)

		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
}
