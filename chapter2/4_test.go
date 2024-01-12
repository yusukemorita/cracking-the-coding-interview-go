package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestPartition(t *testing.T) {
	linkedList := &GenericLinkedListNode[int]{
		value: 3,
		pointer: &GenericLinkedListNode[int]{
			value: 1,
			pointer: &GenericLinkedListNode[int]{
				value: 0,
				pointer: &GenericLinkedListNode[int]{
					value: 3,
				},
			},
		},
	}

	Partition(linkedList, 2)

	if !reflect.DeepEqual(linkedList.values(), []int{0, 1, 3, 3}) && !reflect.DeepEqual(linkedList.values(), []int{1, 0, 3, 3}) {
		t.Errorf("not partitioned correctly. %v", linkedList.values())
	}
}

func TestPartition_Border(t *testing.T) {
	linkedList := &GenericLinkedListNode[int]{
		value: 2,
		pointer: &GenericLinkedListNode[int]{
			value: 3,
			pointer: &GenericLinkedListNode[int]{
				value: 1,
				pointer: &GenericLinkedListNode[int]{
					value: 0,
					pointer: &GenericLinkedListNode[int]{
						value: 3,
					},
				},
			},
		},
	}

	Partition(linkedList, 2)

	// first part should have 0 and 1
	firstPart := []int{}
	firstPart = append(firstPart, linkedList.value)
	firstPart = append(firstPart, linkedList.pointer.value)

	slices.Sort[[]int](firstPart)
	if !reflect.DeepEqual(firstPart, []int{0, 1}) {
		t.Error("first part should contain 0 and 1")
	}

	// second part should have 2, 3, 3
	secondPart := linkedList.pointer.pointer.values()
	slices.Sort[[]int](secondPart)
	if !reflect.DeepEqual(secondPart, []int{2, 3, 3}) {
		t.Error("second part should contain 2, 3, 3")
	}
}
