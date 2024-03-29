package main

import (
	"fmt"
	"testing"
)

func TestGetRandomNode(t *testing.T) {
	t.Run("returns root when there is only one node", func(t *testing.T) {
		root := &RandomNode{value: 15, size: 1}
		tree := RandomTree{root: root}

		random := tree.GetRandomNode()
		if random.value != 15 {
			t.Error()
		}
	})

	t.Run("returns a node in the tree", func(t *testing.T) {
		//   15
		//  13 20
		root := &RandomNode{value: 15, size: 3}
		right := &RandomNode{value: 20, size: 1}
		root.right = right
		left := &RandomNode{value: 13, size: 1}
		root.left = left

		tree := RandomTree{root: root}

		counts := make(map[int]int)
		for i := 0; i <= 100000; i++ {
			random := tree.GetRandomNode()
			counts[random.value]++
		}

		for value, count := range counts {
			fmt.Printf("%d: %d times\n", value, count)
		}
	})
}
