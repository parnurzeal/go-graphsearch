package main

import "fmt"

type Graph struct {
	nodes map[int]Node
}

type Node struct {
	state   string
	adjNode map[int]struct{}
}

func (g *Graph) addEdge(src, dst int) {
	// expand if no this node in graph
	if src < 0 || dst < 0 {
		panic("Node number can't be below 0")
	}
	// make new node if doesn't have one
	if _, ok := g.nodes[src]; !ok {
		g.nodes[src] = Node{adjNode: make(map[int]struct{})}
	}
	if _, ok := g.nodes[dst]; !ok {
		g.nodes[dst] = Node{adjNode: make(map[int]struct{})}
	}
	// connect by adjacent list
	g.nodes[src].adjNode[dst] = struct{}{}
	g.nodes[dst].adjNode[src] = struct{}{}
}

func main() {
	a := Graph{nodes: make(map[int]Node)}
	a.addEdge(0, 1)
	a.addEdge(2, 3)
	a.addEdge(0, 2)
	fmt.Println(a)
}
