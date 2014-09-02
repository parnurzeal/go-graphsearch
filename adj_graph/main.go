package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Graph struct {
	nodes map[int]Node
}

type Node struct {
	payload  []byte
	adjNodes map[int]struct{}
}

func (g *Graph) addNode(key int, data []byte) {
	if _, ok := g.nodes[key]; ok {
		return
	}
	g.nodes[key] = Node{payload: data, adjNodes: make(map[int]struct{})}
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

func (m *Maze) createMazeGraph(in []byte) {
	g := Graph{nodes: make(map[int]Node)}
	// create matrix for easy translating to graph
	matrix := [][]byte{}
	line := []byte{}
	for _, char := range in {
		if char != []byte("\n")[0] {
			line = append(line, char)
		} else {
			matrix = append(matrix, line)
			line = []byte{}
		}
	}
	// create graph from matrix
	count := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			g.addNode(count, []byte{matrix[r][c]})
			count++
		}
	}
	// connect edge
	count = 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			// left
			if c != 0 {
				g.addEdge(count, count-1)
			}
			// right
			if c < len(matrix[r])-1 {
				g.addEdge(count, count+1)
			}
			// up
			if r != 0 && c < len(matrix[r-1]) {
				dst := count - len(matrix[r-1])
				g.addEdge(count, dst)
			}
			// down
			if r < len(matrix)-1 && c < len(matrix[r+1]) {
				dst := count + len(matrix[r])
				g.addEdge(count, dst)
			}
			// next node
			count++
		}
	}
	m.g = g
}

type Maze struct {
	g       Graph
	passed  []bool
	answers []int
}

func (m *Maze) findPathDfs() {
	// set all passed to false
	m.passed = make([]bool, len(m.g.nodes))
	// find S
	for idx, n := range m.g.nodes {
		if string(n.payload) == "S" {
			m.dfs(idx)
		}
	}
	return
}

func (m *Maze) dfs(cur int) bool {
	if _, ok := m.g.nodes[cur]; !ok {
		return false
	}
	if m.passed[cur] || string(m.g.nodes[cur].payload) == "|" || string(m.g.nodes[cur].payload) == "-" {
		return false
	}
	m.passed[cur] = true
	if string(m.g.nodes[cur].payload) == "E" {
		m.answers = append(m.answers, cur)
		fmt.Println("Found")
		return true
	}
	for adj, _ := range m.g.nodes[cur].adjNodes {
		if m.dfs(adj) {
			m.answers = append(m.answers, cur)
			return true
		}
	}
	return false
}

func mimicAnswers(in []byte, answers []int) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	count := 0
	for idx, char := range in {
		for _, tmp := range answers {
			if tmp == count {
				out[idx] = []byte("X")[0]
			}
		}
		if char != []byte("\n")[0] {
			count++
		}
	}
	return out
}

func main() {
	inBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(inBytes))
	m := Maze{answers: make([]int, 0)}
	m.createMazeGraph(inBytes)
	m.findPathDfs()
	fmt.Println(string(mimicAnswers(inBytes, m.answers)))

}
