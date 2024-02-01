package main

// Implement bidirectional (breadth first) search, as it is more efficient
// than breadth first searching from one node.
// Assume all nodes have a unique value. If there is no unique value that
// can be used like an id, then we could add an `isVisited` boolean value
// to the nodes, and do a search from one node instead of a bidirectional search.
func HasRouteBetweenNodes(nodeA, nodeB GraphNode[string]) bool {
	nodeAVisited := make(map[string]bool)
	nodeBVisited := make(map[string]bool)

	nodeANext := []*GraphNode[string]{
		&nodeA,
	}
	nodeBNext := []*GraphNode[string]{
		&nodeB,
	}

	hasRoute := false

outer:
	for {
		if len(nodeANext) == 0 && len(nodeBNext) == 0 {
			break
		}

		// A: visit all nodes in next level
		var newNodeANext []*GraphNode[string]
		for _, node := range nodeANext {
			newNodeANext = append(newNodeANext, node.children...)
			// if the value is in nodeBVisited, there is a route
			if _, bVisited := nodeBVisited[node.value]; bVisited {
				hasRoute = true
				break outer
			}

			nodeAVisited[node.value] = true
		}
		nodeANext = newNodeANext

		// B: visit all nodes in next level
		var newNodeBNext []*GraphNode[string]
		for _, node := range nodeBNext {
			newNodeBNext = append(newNodeBNext, node.children...)
			// if the value is in nodeAVisited, there is a route
			if _, aVisited := nodeAVisited[node.value]; aVisited {
				hasRoute = true
				break outer
			}

			nodeBVisited[node.value] = true
		}
		nodeBNext = newNodeBNext
	}

	return hasRoute
}

type GraphNode[T comparable] struct {
	value    T
	children []*GraphNode[T]
}
