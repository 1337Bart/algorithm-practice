package main

import (
	"fmt"

	"github.com/1337Bart/algorithm-practice/heap"
)

type Edge struct {
	To     int
	Weight int
}

type Graph [][]Edge

func Dijkstra(g Graph, source int) []int {
	// Initialize the shortest paths with infinity
	INF := 1 << 60
	dist := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		dist[i] = INF
	}
	dist[source] = 0

	// Initialize the priority queue
	h := heap.NewBinaryHeap()
	h.Insert(&heap.Item{Node: source, Weight: 0})

	for len(h.Heap) > 0 {
		// Get the node with the smallest weight
		item := h.DeleteMin()
		v := item.Node

		// Ensure we don't process a node more than once
		if dist[v] < item.Weight {
			continue
		}

		// Examine the edges leading to the neighbors of the current node
		for _, edge := range g[v] {
			// If we found a shorter path to the neighbor node, update it
			if newDist := dist[v] + edge.Weight; newDist < dist[edge.To] {
				dist[edge.To] = newDist
				h.Insert(&heap.Item{Node: edge.To, Weight: newDist})
			}
		}
	}

	return dist
}

func main() {
	// Define a graph - each key is a node and each value is a dictionary of the neighboring nodes and the weights of the edges to them
	g := Graph{
		[]Edge{{To: 1, Weight: 2}, {To: 2, Weight: 5}, {To: 4, Weight: 10}},
		[]Edge{{To: 2, Weight: 4}, {To: 3, Weight: 6}},
		[]Edge{{To: 1, Weight: 1}, {To: 3, Weight: 3}},
		[]Edge{{To: 4, Weight: 2}},
		[]Edge{},
	}

	dist := Dijkstra(g, 0)

	// Print the shortest paths
	fmt.Println(dist)
}
