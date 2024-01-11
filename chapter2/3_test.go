package main

import (
	"reflect"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	middleNode := &LinkedListNode{
		value: "2",
		pointer: &LinkedListNode{
			value: "3",
			pointer: &LinkedListNode{
				value: "4",
			},
		},
	}

	linkedList := &LinkedListNode{
		value: "0",
		pointer: &LinkedListNode{
			value:   "1",
			pointer: middleNode,
		},
	}

	DeleteMiddleNode(middleNode)

	got := linkedList.values()
	expected := []string{"0", "1", "3", "4"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
