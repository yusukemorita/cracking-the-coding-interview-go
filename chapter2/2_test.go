package main

import (
	"fmt"
	"testing"
)

func TestKthToLast(t *testing.T) {
	functions := map[string](func(head *LinkedListNode, k int) (*LinkedListNode, error)){
		"KthToLast1": KthToLast1,
	}

	for name, function := range functions {
		t.Run(name, func(t *testing.T) {

			t.Run("returns kth to last node", func(t *testing.T) {
				linkedList := &LinkedListNode{
					value: "0",
					pointer: &LinkedListNode{
						value: "1",
						pointer: &LinkedListNode{
							value: "2",
							pointer: &LinkedListNode{
								value: "3",
								pointer: &LinkedListNode{
									value: "4",
									pointer: &LinkedListNode{
										value: "5",
										pointer: &LinkedListNode{
											value: "6",
										},
									},
								},
							},
						},
					},
				}

				got, err := function(linkedList, 2)
				if err != nil {
					fmt.Print(err)
					t.Error("expected no error")
				}
				if got.value != "5" {
					t.Errorf("expected 2nd to last, got %s", got.value)
				}
			})

			t.Run("returns error when k is 0", func(t *testing.T) {
				linkedList := &LinkedListNode{
					value: "0",
					pointer: &LinkedListNode{
						value: "1",
						pointer: &LinkedListNode{
							value: "2",
							pointer: &LinkedListNode{
								value: "3",
								pointer: &LinkedListNode{
									value: "4",
									pointer: &LinkedListNode{
										value: "5",
										pointer: &LinkedListNode{
											value: "6",
										},
									},
								},
							},
						},
					},
				}

				_, err := function(linkedList, 0)

				if err == nil {
					t.Error("expected error")
				}
			})

			t.Run("returns error when k is length + 1", func(t *testing.T) {
				linkedList := &LinkedListNode{
					value: "0",
					pointer: &LinkedListNode{
						value: "1",
						pointer: &LinkedListNode{
							value: "2",
							pointer: &LinkedListNode{
								value: "3",
								pointer: &LinkedListNode{
									value: "4",
									pointer: &LinkedListNode{
										value: "5",
										pointer: &LinkedListNode{
											value: "6",
										},
									},
								},
							},
						},
					},
				}

				_, err := function(linkedList, 8)

				if err == nil {
					t.Error("expected error")
				}
			})
		})
	}
}

func TestKthToLast2(t *testing.T) {
	t.Run("returns kth to last node", func(t *testing.T) {
		linkedList := &LinkedListNode{
			value: "0",
			pointer: &LinkedListNode{
				value: "1",
				pointer: &LinkedListNode{
					value: "2",
					pointer: &LinkedListNode{
						value: "3",
						pointer: &LinkedListNode{
							value: "4",
							pointer: &LinkedListNode{
								value: "5",
								pointer: &LinkedListNode{
									value: "6",
								},
							},
						},
					},
				},
			},
		}

		counter := Counter{0}
		got := KthToLast2(linkedList, 2, &counter)
		if got.value != "5" {
			t.Errorf("expected 2nd to last, got %s", got.value)
		}
	})
}
