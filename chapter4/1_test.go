package main

import "testing"

func TestHasRouteBetweenNodes(t *testing.T) {
	t.Run("returns false when there is no route", func(t *testing.T) {
		nodeA := GraphNode[string]{
			value: "A",
			children: []*GraphNode[string]{
				{
					value: "C",
				},
				{
					value: "D",
				},
			},
		}

		nodeB := GraphNode[string]{
			value: "B",
		}

		got := HasRouteBetweenNodes(nodeA, nodeB)
		if got {
			t.Error("expected false")
		}
	})

	t.Run("returns true when there is a route", func(t *testing.T) {
		nodeB := GraphNode[string]{
			value: "B",
		}

		nodeA := GraphNode[string]{
			value: "A",
			children: []*GraphNode[string]{
				{
					value: "C",
				},
				{
					value: "D",
				},
				&nodeB,
			},
		}

		got := HasRouteBetweenNodes(nodeA, nodeB)
		if !got {
			t.Error("expected true")
		}
	})
}
