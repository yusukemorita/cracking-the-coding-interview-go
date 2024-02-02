package main

import "fmt"

type dependency struct {
	dependant string
	depended  string
}

type Project struct {
	name            string
	dependencyCount int
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
