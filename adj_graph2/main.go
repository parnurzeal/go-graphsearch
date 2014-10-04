package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Graph struct {
	nodes map[int]Node
}

type Node struct {
	payload byte
	adj     map[int]struct{}
}

func NewGraph() Graph {
	return Graph{make(map[int]Node)}
}

func (g *Graph) addNode(key int, data byte) {
	g.nodes[key] = Node{payload: data, adj: make(map[int]struct{})}
}

func (g *Graph) addEdge(src, dst int) {
	g.nodes[src].adj[dst] = struct{}{}
	g.nodes[dst].adj[src] = struct{}{}
}

func createMazeGraph(in []byte) Graph {
	g := NewGraph()
	// create matrix
	matrix := make([][]byte, 0)
	line := make([]byte, 0)
	for _, block := range in {
		if block != '\n' {
			line = append(line, block)
		} else {
			matrix = append(matrix, line)
			line = make([]byte, 0)
		}
		fmt.Print(string(block))
	}
	// create node
	count := 0
	for _, row := range matrix {
		for _, col := range row {
			g.addNode(count, col)
			count++
		}
	}
	// create edge
	length := len(matrix[0])
	for r, row := range matrix {
		for c, _ := range row {
			if c-1 >= 0 {
				g.addEdge(r*length+c, r*length+c-1)
			}
			if r-1 >= 0 {
				g.addEdge(r*length+c, (r-1)*length+c)
			}
			if r+1 < len(matrix) {
				g.addEdge(r*length+c, (r+1)*length+c)
			}
			if c+1 < len(row) {
				g.addEdge(r*length+c, r*length+c+1)
			}
		}
	}
	return g
}

func (g *Graph) bfs() []int {
	sNode, err := g.findStartNode()
	if err != nil {
		panic(err)
	}
	visited := make(map[int]struct{})
	parent := make(map[int]int)
	queue := []int{sNode}
	var answer []int
	for len(queue) != 0 {
		cur := queue[0]
		visited[cur] = struct{}{}
		queue = queue[1:]
		switch g.nodes[cur].payload {
		case ' ', 'S':
			for key, _ := range g.nodes[cur].adj {
				if _, ok := visited[key]; ok {
					continue
				}
				parent[key] = cur
				queue = append(queue, key)
			}
		case '|', '-':
			continue
		case 'E':
			answer = append(answer, cur)
			// backtrack to get answer
			for _, ok := parent[cur]; ok; _, ok = parent[cur] {
				cur = parent[cur]
				answer = append(answer, cur)
			}
			break
		}
	}
	return answer
}

func (g *Graph) findStartNode() (int, error) {
	for key, node := range g.nodes {
		if node.payload == 'S' {
			return key, nil
		}
	}
	return 0, errors.New("Cannot find start node")
}

func (g *Graph) dfs() []int {
	sNode, err := g.findStartNode()
	if err != nil {
		panic(err)
	}
	visited := make(map[int]struct{})
	answer := make([]int, 0)
	g.dfs2(sNode, visited, &answer)
	return answer
}

func (g *Graph) dfs2(cur int, visited map[int]struct{}, answer *[]int) bool {
	if _, ok := visited[cur]; ok {
		return false
	}
	visited[cur] = struct{}{}
	switch g.nodes[cur].payload {
	case '|', '-':
		return false
	case 'E':
		return true
	}
	for i, _ := range g.nodes[cur].adj {
		if g.dfs2(i, visited, answer) {
			*answer = append(*answer, i)
			return true
		}
	}
	return false
}

func mimicAnswer(in []byte, answer []int) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	count := 0
	for idx, char := range in {
		for _, pos := range answer {
			if count == pos {
				out[idx] = 'X'
				break
			}
		}
		if char != '\n' {
			count++
		}
	}
	return out
}

func main() {
	in, _ := ioutil.ReadAll(os.Stdin)
	g := createMazeGraph(in)
	answer := g.dfs()
	out := mimicAnswer(in, answer)
	fmt.Println(string(out))
	answer = g.bfs()
	out = mimicAnswer(in, answer)
	fmt.Println(string(out))
}
