package main

import (
	"fmt"
	"testing"
)

// difficult to test, so just print the BST
func TestBuildDepthLists(t *testing.T) {
	t.Run("", func(t *testing.T) {
		bst := BSTNode{
			value: 5,
			left: &BSTNode{
				value: 3,
				left: &BSTNode{
					value: 2,
				},
			},
			right: &BSTNode{
				value: 7,
				left: &BSTNode{
					value: 6,
				},
				right: &BSTNode{
					value: 8,
				},
			},
		}

		linkedLists := BuildDepthLists(&bst)
		if len(linkedLists) != 3 {
			t.Fatal("expected 3 lists")
		}

		fmt.Println(linkedLists[0].values())
		fmt.Println(linkedLists[1].values())
		fmt.Println(linkedLists[2].values())
	})
}
