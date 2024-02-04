package main

import "math/rand"

type RandomNode struct {
	value int
	size  int
	left  *RandomNode
	right *RandomNode
}

type RandomTree struct {
	root     *RandomNode
	allNodes map[int][]*RandomNode
}

// solution 1. maintain an slice of all nodes in BinaryTree struct
// - Find: would not affect the runtime
// - Insert: adding nodes to the slice can be done in (amortized) constant time
// - Delete: deleting nodes would be done in O(N) time.
//           If this is just a binary tree, then Delete would also take O(N) time,
//           however if this is a BST, then finding the node in the tree would only take O(N)
//           time, so this would slow down the Delete operation.
// - GetRandomNode: pick a random number in 0 ~ len(slice)-1, and return the item at the index.
//                  runtime is constant.

// solution 2. maintain a hashmap of all nodes, key: node value, value: slice of all nodes with the same value
// - Find: would make Find faster
// - Insert: adding nodes to the map can be done in O(x) time, where x is the number of nodes with the same value
// - Delete: same as Insert
// - GetRandomNode: Extract all nodes from the map into an array, and then do as solution 1.
//                  extracting the nodes will take O(N) time.

// solution 3. pick a random number to decide to traverse left, right or use the current node value
//             balance the probabilities, so the size of the subtree is considered
// - Find: unchanged
// - Insert: increment the size of all parents of the inserted node.
//           runtime complexity should stay the same.
// - Delete: unchanged

func (tree *RandomTree) GetRandomNode() *RandomNode {
	return getRandomNode(tree.root)
}

func getRandomNode(node *RandomNode) *RandomNode {
	if node.left == nil && node.right == nil {
		return node
	}

	weightNodes := []WeightNode{
		{
			node:   node,
			weight: 1,
		},
	}

	if node.left != nil {
		weightNodes = append(weightNodes, WeightNode{
			node:   node.left,
			weight: node.left.size,
		})
	}

	if node.right != nil {
		weightNodes = append(weightNodes, WeightNode{
			node:   node.right,
			weight: node.right.size,
		})
	}

	return weightedRandomSelect(weightNodes)
}

type WeightNode struct {
	node        *RandomNode
	weight      int
	randomRange []int // tuple of two values, inclusive range
}

func weightedRandomSelect(weightNodes []WeightNode) *RandomNode {
	var weighted []*RandomNode

	for _, wn := range weightNodes {
		for i := 1; i <= wn.weight; i++ {
			weighted = append(weighted, wn.node)
		}
	}

	randomIndex := rand.Intn(len(weighted))
	return weighted[randomIndex]
}
