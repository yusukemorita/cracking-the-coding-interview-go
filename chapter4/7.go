package main

import "fmt"

type dependency struct {
	dependant string
	depended  string
}

type Project struct {
	name            string
	dependencyCount int
	hasBeenOrdered  bool
}

func BuildOrder(projectNames []string, dependencies []dependency) ([]string, bool) {
	allNodes := buildGraphNodes(projectNames, dependencies)

	var queue []GraphNode[Project]

	// find "root" nodes (nodes that don't have any dependencies) and
	// add them to the queue
	for _, node := range allNodes {
		if node.value.dependencyCount == 0 {
			queue = append(queue, node)
		}
	}

	var order []string
	for {
		if len(queue) == 0 {
			break
		}

		current := queue[0]
		queue = queue[1:]

		order = append(order, current.value.name)

		for _, child := range current.children {
			child.value.dependencyCount--
			if child.value.dependencyCount == 0 {
				queue = append(queue, *child)
			}
		}
	}

	if len(order) != len(projectNames) {
		return nil, false
	}

	return order, true
}

// children: nodes that must be built AFTER the node
func buildGraphNodes(projectNames []string, dependencies []dependency) []GraphNode[Project] {
	allNodes := make(map[string]*GraphNode[Project])
	for _, name := range projectNames {
		allNodes[name] = &GraphNode[Project]{
			value: Project{
				name:            name,
				dependencyCount: 0,
			},
			children: []*GraphNode[Project]{},
		}
	}

	for _, d := range dependencies {
		dependantNode, ok := allNodes[d.dependant]
		if !ok {
			panic(fmt.Sprintf("node %s not found", d.dependant))
		}

		dependedNode, ok := allNodes[d.depended]
		if !ok {
			panic(fmt.Sprintf("node %s not found", d.depended))
		}

		dependedNode.children = append(dependedNode.children, dependantNode)
		dependantNode.value.dependencyCount++
	}

	var slice []GraphNode[Project]
	for _, node := range allNodes {
		slice = append(slice, *node)
	}

	return slice
}

func BuildOrder2(projectNames []string, dependencies []dependency) ([]string, bool) {
	allNodes := buildGraphNodes2(projectNames, dependencies)

	var ordered []string

	for {
		var visited []string
		unorderedNode := findFirstUnorderedNode(allNodes)
		if unorderedNode == nil {
			// finished ordering all nodes
			break
		}
		fmt.Printf("unordered node name: %s\n", unorderedNode.value.name)

		hasCycle := depthFirstSearch(unorderedNode, &ordered, &visited)
		if hasCycle {
			return nil, false
		}
	}

	return ordered, true
}

func contains[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}

	return false
}

func depthFirstSearch(node *GraphNode[Project], ordered, visited *[]string) (hasCycle bool) {
	fmt.Printf("visiting %s\n", node.value.name)

	if node.value.hasBeenOrdered {
		fmt.Printf("project %s has already been ordered\n", node.value.name)
		return false
	}

	if contains(*visited, node.value.name) {
		fmt.Printf("project %s has already been visited\n", node.value.name)
		return true
	}

	*visited = append(*visited, node.value.name)

	for _, child := range unorderedChildren(node) {
		childHasCycle := depthFirstSearch(child, ordered, visited)
		if childHasCycle {
			return true
		}
	}

	if len(unorderedChildren(node)) == 0 {
		fmt.Printf("adding %s to order\n", node.value.name)
		node.value.hasBeenOrdered = true
		*ordered = append(*ordered, node.value.name)
		return false
	}

	return false
}

func findFirstUnorderedNode(nodes []*GraphNode[Project]) *GraphNode[Project] {
	for _, node := range nodes {
		if !node.value.hasBeenOrdered {
			return node
		}
	}

	return nil
}

// children: nodes that must be built BEFORE the node
func buildGraphNodes2(projectNames []string, dependencies []dependency) []*GraphNode[Project] {
	allNodes := make(map[string]*GraphNode[Project])
	for _, name := range projectNames {
		allNodes[name] = &GraphNode[Project]{
			value: Project{
				name:            name,
				dependencyCount: 0,
			},
			children: []*GraphNode[Project]{},
		}
	}

	for _, d := range dependencies {
		dependantNode, ok := allNodes[d.dependant]
		if !ok {
			panic(fmt.Sprintf("node %s not found", d.dependant))
		}

		dependedNode, ok := allNodes[d.depended]
		if !ok {
			panic(fmt.Sprintf("node %s not found", d.depended))
		}

		dependantNode.children = append(dependantNode.children, dependedNode)
		// dependedNode.value.dependencyCount++
	}

	var slice []*GraphNode[Project]
	for _, node := range allNodes {
		slice = append(slice, node)
	}

	return slice
}

func unorderedChildren(node *GraphNode[Project]) []*GraphNode[Project] {
	var result []*GraphNode[Project]
	var names []string
	for _, child := range node.children {
		if !child.value.hasBeenOrdered {
			result = append(result, child)
			names = append(names, child.value.name)
		}
	}
	fmt.Printf("unordered child names: %v\n", names)
	return result
}
