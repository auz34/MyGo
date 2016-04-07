// Given a directed graph, design an algorithm to find out whether there is a route between
// two nodes.

package main

import (
	"fmt"
	"cracking/ch4/basics"
)

func routeExist(graph *basics.DirectedGraph, from, to int) bool {
	visited := make(map[int]bool)

	var traverse func (nodeId int) bool
	traverse = func (nodeId int) bool {
		if nodeId == to {
			return true
		}

		visited[nodeId] = true
		v := graph.Vertices[nodeId]

		for toId, _ := range v.ToEdges {
			if traverse(toId) {
				return true
			}
		}

		return false
	}

	return traverse(from)
}

func main() {
	graph := basics.GenerateDirectedGraphSample()

	fmt.Printf("There is path from #1 to #12: %v \n", routeExist(graph, 1, 12))
	fmt.Printf("There is path from #1 to #8: %v \n", routeExist(graph, 1, 8))
	fmt.Printf("There is path from #9 to #12: %v \n", routeExist(graph, 9, 12))
}
