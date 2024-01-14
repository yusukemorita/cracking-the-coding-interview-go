package main

import (
	"testing"
)

func TestSetOfStacks(t *testing.T) {
	setOfStacks := SetOfStacks{
		stackCapacity: 3,
	}

	setOfStacks.push("1")
	setOfStacks.push("2")
	setOfStacks.push("3")
	setOfStacks.push("4")
	setOfStacks.push("5")

	popped, ok := setOfStacks.pop()
	if !ok || popped != "5" {
		t.Error()
	}

	popped, ok = setOfStacks.pop()
	if !ok || popped != "4" {
		t.Error()
	}

	popped, ok = setOfStacks.pop()
	if !ok || popped != "3" {
		t.Error()
	}

	popped, ok = setOfStacks.pop()
	if !ok || popped != "2" {
		t.Error()
	}

	popped, ok = setOfStacks.pop()
	if !ok || popped != "1" {
		t.Error()
	}

	popped, ok = setOfStacks.pop()
	if ok {
		t.Error()
	}
}

func TestSetOfStacks_Push(t *testing.T) {
	t.Run("creates a new stack when no existing stacks", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 10,
		}

		setOfStacks.push("abc")

		if len(setOfStacks.stacks) != 1 {
			t.Errorf("expected length of stacks to be 1, got %d", len(setOfStacks.stacks))
		}
	})

	t.Run("creates a new stack when last existing stack is full", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stacks: []*Stack{
				{
					length: 5,
					head: &LinkedListNode[string]{
						value: "1",
						next: &LinkedListNode[string]{
							value: "2",
							next: &LinkedListNode[string]{
								value: "3",
								next: &LinkedListNode[string]{
									value: "4",
									next: &LinkedListNode[string]{
										value: "5",
									},
								},
							},
						},
					},
				},
			},
			stackCapacity: 5,
		}

		setOfStacks.push("abc")

		if len(setOfStacks.stacks) != 2 {
			t.Errorf("expected length of stacks to be 2, got %d", len(setOfStacks.stacks))
		}
	})
}

func TestSetOfStacks_Pop(t *testing.T) {
	t.Run("returns the last item in the last stack", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 3,
			stacks: []*Stack{
				{
					length: 3,
					head: &LinkedListNode[string]{
						value: "3",
						next: &LinkedListNode[string]{
							value: "2",
							next: &LinkedListNode[string]{
								value: "1",
							},
						},
					},
				},
				{
					length: 2,
					head: &LinkedListNode[string]{
						value: "5",
						next: &LinkedListNode[string]{
							value: "4",
						},
					},
				},
			},
		}

		got, ok := setOfStacks.pop()

		if ok != true {
			t.Error("expected pop to succeed")
		}

		if got != "5" {
			t.Errorf("expected pop to return last item, got %s", got)
		}
	})

	t.Run("removes the last stack if it becomes empty", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 3,
			stacks: []*Stack{
				{
					length: 3,
					head: &LinkedListNode[string]{
						value: "3",
						next: &LinkedListNode[string]{
							value: "2",
							next: &LinkedListNode[string]{
								value: "1",
							},
						},
					},
				},
				{
					length: 1,
					head: &LinkedListNode[string]{
						value: "4",
					},
				},
			},
		}

		got, ok := setOfStacks.pop()

		if ok != true {
			t.Error("expected pop to succeed")
		}

		if got != "4" {
			t.Errorf("expected pop to return last item, got %s", got)
		}

		if len(setOfStacks.stacks) != 1 {
			t.Errorf("expected empty stack to be removed, length: %d", len(setOfStacks.stacks))
		}
	})

	t.Run("returns false when no items in stacks", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 3,
			stacks:        []*Stack{},
		}

		_, ok := setOfStacks.pop()

		if ok {
			t.Errorf("expected ok to be false")
		}
	})
}

func TestSetOfStacks_PopSubStack(t *testing.T) {
	t.Run("returns the last item in the specified substack", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 3,
			stacks: []*Stack{
				{
					length: 3,
					head: &LinkedListNode[string]{
						value: "3",
						next: &LinkedListNode[string]{
							value: "2",
							next: &LinkedListNode[string]{
								value: "1",
							},
						},
					},
				},
				{
					length: 2,
					head: &LinkedListNode[string]{
						value: "5",
						next: &LinkedListNode[string]{
							value: "4",
						},
					},
				},
			},
		}

		got, ok := setOfStacks.popSubStack(0)

		if ok != true {
			t.Error("expected pop to succeed")
		}

		if got != "3" {
			t.Errorf("expected pop to return last item, got %s", got)
		}
	})

	t.Run("removes the stack if it becomes empty", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 1,
			stacks: []*Stack{
				{
					length: 1,
					head: &LinkedListNode[string]{
						value: "1",
					},
				},
				{
					length: 1,
					head: &LinkedListNode[string]{
						value: "2",
					},
				},
			},
		}

		got, ok := setOfStacks.popSubStack(0)

		if ok != true {
			t.Error("expected pop to succeed")
		}

		if got != "1" {
			t.Errorf("expected pop to return last item of specified substack, got %s", got)
		}

		if len(setOfStacks.stacks) != 1 {
			t.Errorf("expected empty stack to be removed, length: %d", len(setOfStacks.stacks))
		}
	})

	// to implement this operation, we would need a doubly linked list
	// skip for now
	t.Run("shifts items in substacks so that each substack except the last is at full capacity", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 3,
			stacks: []*Stack{
				{
					length: 3,
					head: &LinkedListNode[string]{
						value: "3",
						next: &LinkedListNode[string]{
							value: "2",
							next: &LinkedListNode[string]{
								value: "1",
							},
						},
					},
				},
				{
					length: 3,
					head: &LinkedListNode[string]{
						value: "6",
						next: &LinkedListNode[string]{
							value: "5",
							next: &LinkedListNode[string]{
								value: "4",
							},
						},
					},
				},
				{
					length: 1,
					head: &LinkedListNode[string]{
						value: "7",
					},
				},
			},
		}

		popped, ok := setOfStacks.popSubStack(0)
		if !ok || popped != "3" {
			t.Errorf("expected ok: true, popped: 3, got ok: %v, popped: %s", ok, popped)
		}

		if len(setOfStacks.stacks) != 2 {
			t.Errorf("expected empty stacks to be removed")
		}

		if setOfStacks.stacks[0].length != 3 || setOfStacks.stacks[1].length != 3 {
			t.Errorf("expected remaining stacks to be at full capacity")
		}
	})

	t.Run("returns false when no items in stacks", func(t *testing.T) {
		setOfStacks := SetOfStacks{
			stackCapacity: 3,
			stacks:        []*Stack{},
		}

		_, ok := setOfStacks.pop()

		if ok {
			t.Errorf("expected ok to be false")
		}
	})
}
