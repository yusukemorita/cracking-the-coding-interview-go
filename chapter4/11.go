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

// solution 3. pick one random number X, and return the Xth node in in order traversal.
// - Find: unchanged
// - Insert: increment the size of all parents of the inserted node.
//           runtime complexity should stay the same.
// - Delete: unchanged
// - GetRandomNode: time complexity is O(logN) if balanced, or O(D) (D = depth)

func (tree *RandomTree) GetRandomNode() *RandomNode {
	random := rand.Intn(tree.root.size)

	return getRandomNode(random, tree.root)
}

func getRandomNode(random int, node *RandomNode) *RandomNode {
	// when left size is L, and right size is R
	// random = 0               ... node
	// 1 <= random <= L         ... left
	// L + 1 <= random <= L + R ... right
	leftSize := 0
	if node.left != nil {
		leftSize = node.left.size
	}

	if random == 0 {
		return node
	}

	// 1 <= random <= L
	if 1 <= random && random <= leftSize {
		return getRandomNode(random-1, node.left)
	}

	// L + 1 <= random <= L + R
	return getRandomNode(random-1-leftSize, node.right)
}
