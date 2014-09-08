package main

type Graph struct {
	nodes map[int]Node
}

type Node struct {
	payload  Interface
	adjNodes map[int]struct{}
}

func (g *Graph) addNode(key int, payload Interface) {
	if _, ok := g.nodes[key]; ok {
		return
	}
	g.nodes[key] = Node{payload: payload, adjNodes: make(map[int]struct{})}
}

func (g *Graph) addEdge(src, dst int) {
	// expand if no this node in graph
	if src < 0 || dst < 0 {
		panic("Node number can't be below 0")
	}
	// make new node if doesn't have one
	if _, ok := g.nodes[src]; !ok {
		panic("No node " + string(src))
	}
	if _, ok := g.nodes[dst]; !ok {
		panic("No node " + string(dst))
	}
	// connect by adjacent list
	g.nodes[src].adjNodes[dst] = struct{}{}
	g.nodes[dst].adjNodes[src] = struct{}{}
}

type Interface interface {
	Compare() State
}

type State int

const (
	START State = iota
	END
	SPACE
	OBSTACLE
)

func (g *Graph) bfsFunc() {
	// find start node
	/*startIdx := g.findStartNode()
	// declare parent for backtrace solution
	parent := make([]int, len(g.nodes))
	answers := []int{}

	q := Queue{}
	q.enqueue(startIdx)
	for !q.isEmpty() {
		cur := q.dequeue().(int)

	}*/
}

func (g *Graph) bfs() []int {
	passed := make([]bool, len(g.nodes))
	// find start node
	startIdx := g.findStartNode()
	// declare parent for backtrack solution
	parent := make([]int, len(g.nodes))
	answers := []int{}

	q := Queue{}
	q.enqueue(startIdx)
	for !q.isEmpty() {
		cur := q.dequeue().(int)
		if passed[cur] {
			continue
		}
		passed[cur] = true
		if g.nodes[cur].payload.Compare() == OBSTACLE {
			continue
		}
		if g.nodes[cur].payload.Compare() == END {
			// backtrace
			answers = append(answers, cur)
			for tmpIdx := cur; g.nodes[tmpIdx].payload.Compare() != START; tmpIdx = parent[tmpIdx] {
				answers = append(answers, tmpIdx)
			}
			break
		}
		for adjIdx, _ := range g.nodes[cur].adjNodes {
			if !passed[adjIdx] {
				parent[adjIdx] = cur
				q.enqueue(adjIdx)
			}
		}
	}
	return answers
}

func (g *Graph) findStartNode() int {
	for idx, n := range g.nodes {
		if n.payload.Compare() == START {
			return idx
		}
	}
	return -1
}

func (g *Graph) dfs() []int {
	// set []passed to all false
	passed := make([]bool, len(g.nodes))
	// find start node
	startNode := g.findStartNode()
	// declare answers
	answers := []int{}
	g.dfsInner(startNode, &passed, &answers)
	return answers
}

func (g *Graph) dfsInner(cur int, passed *[]bool, answers *[]int) bool {
	if _, ok := g.nodes[cur]; !ok {
		return false
	}
	if (*passed)[cur] || g.nodes[cur].payload.Compare() == OBSTACLE {
		return false
	}
	(*passed)[cur] = true
	if g.nodes[cur].payload.Compare() == END {
		(*answers) = append((*answers), cur)
		return true
	}
	for adjIdx, _ := range g.nodes[cur].adjNodes {
		if g.dfsInner(adjIdx, passed, answers) {
			(*answers) = append((*answers), cur)
			return true
		}
	}
	return false
}
