package main

import (
	"reflect"
	"testing"
)

func TestAddNumbersReversed(t *testing.T) {
	t.Run("adds numbers when the sum of each digit is less than 10", func(t *testing.T) {
		// 123
		numberA := &GenericLinkedListNode[int]{
			value: 3,
			pointer: &GenericLinkedListNode[int]{
				value: 2,
				pointer: &GenericLinkedListNode[int]{
					value: 1,
				},
			},
		}

		// 45
		numberB := &GenericLinkedListNode[int]{
			value: 5,
			pointer: &GenericLinkedListNode[int]{
				value: 4,
			},
		}

		// should be 168
		sum := AddNumbersReversed(numberA, numberB)

		if !reflect.DeepEqual(sum.values(), []int{8, 6, 1}) {
			t.Errorf("expected 168(8 -> 6 -> 1), got %v", sum.values())
		}

	})

	t.Run("adds numbers when the sum of each digit may exceed 10", func(t *testing.T) {
		// 123
		numberA := &GenericLinkedListNode[int]{
			value: 3,
			pointer: &GenericLinkedListNode[int]{
				value: 2,
				pointer: &GenericLinkedListNode[int]{
					value: 1,
				},
			},
		}

		// 49
		numberB := &GenericLinkedListNode[int]{
			value: 9,
			pointer: &GenericLinkedListNode[int]{
				value: 4,
			},
		}

		// should be 172
		sum := AddNumbersReversed(numberA, numberB)

		if !reflect.DeepEqual(sum.values(), []int{2, 7, 1}) {
			t.Errorf("expected 172(2 -> 7 -> 1), got %v", sum.values())
		}
	})
}

func TestAddNumbers(t *testing.T) {
	t.Run("adds numbers when the sum of each digit is less than 10", func(t *testing.T) {
		// 123
		numberA := &GenericLinkedListNode[int]{
			value: 1,
			pointer: &GenericLinkedListNode[int]{
				value: 2,
				pointer: &GenericLinkedListNode[int]{
					value: 3,
				},
			},
		}

		// 45
		numberB := &GenericLinkedListNode[int]{
			value: 4,
			pointer: &GenericLinkedListNode[int]{
				value: 5,
			},
		}

		// should be 168
		sum := AddNumbers(numberA, numberB)

		if !reflect.DeepEqual(sum.values(), []int{1, 6, 8}) {
			t.Errorf("expected 168, got %v", sum.values())
		}

	})

	t.Run("adds numbers when the sum of each digit may exceed 10", func(t *testing.T) {
		// 123
		numberA := &GenericLinkedListNode[int]{
			value: 1,
			pointer: &GenericLinkedListNode[int]{
				value: 2,
				pointer: &GenericLinkedListNode[int]{
					value: 3,
				},
			},
		}

		// 49
		numberB := &GenericLinkedListNode[int]{
			value: 4,
			pointer: &GenericLinkedListNode[int]{
				value: 9,
			},
		}

		// should be 172
		sum := AddNumbers(numberA, numberB)

		if !reflect.DeepEqual(sum.values(), []int{1, 7, 2}) {
			t.Errorf("expected 172, got %v", sum.values())
		}
	})
}
