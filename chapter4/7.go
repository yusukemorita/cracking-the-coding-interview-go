package main

import "fmt"

type dependency struct {
	from string
	to   string
}

func BuildOrder(projects []string, dependencies []dependency) ([]string, bool) {
	dependencyMap := make(map[string][]string)
	for _, d := range dependencies {
		dependencyMap[d.from] = append(dependencyMap[d.from], d.to)
	}

	var order []string
	validOrder := false

	for {
		if len(projects) == 0 {
			validOrder = true
			break
		}

		var newProjects []string
		noDependencyProjectFound := false

		for _, project := range projects {
			fmt.Printf("inspecting project %s\n", project)
			deps, ok := dependencyMap[project]
			if !ok || len(deps) == 0 {
				// no dependencies
				order = append(order, project)
				removeDependency(project, dependencyMap)
				noDependencyProjectFound = true
			} else {
				newProjects = append(newProjects, project)
			}
		}

		// if everything has a dependency, then there is no valid build cycle
		if !noDependencyProjectFound {
			break
		}






		projects = newProjects
	}

	return order, validOrder
}


func removeDependency(project string, dependencyMap map[string][]string) {
	for key, values := range dependencyMap {
		dependencyMap[key] = removeValue(values, project)
	}
}

func removeValue(slice []string, value string) []string {
	var new []string
	for _, val := range slice {
		if val != value {
			new = append(new, val)
		}
	}

	return new
}

// func hasNoDependency(project string, dependenci)

// var allNodes map[string]GraphNode[string]
// for _, project := range projects {
// 	allNodes[project] = GraphNode[string]{
// 		value: project,
// 	}
// }

// for _, d := range dependencies {
// 	fromNode, ok := allNodes[d.from]
// 	if !ok {
// 		panic(fmt.Sprintf("node %s not found", d.from))
// 	}

// 	toNode, ok := allNodes[d.to]
// 	if !ok {
// 		panic(fmt.Sprintf("node %s not found", d.to))
// 	}

// 	fromNode.children = append(fromNode.children, &toNode)
// }
