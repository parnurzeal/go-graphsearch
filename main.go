package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Maze struct {
	node    [][]Node
	answers []*Node
}

type Node struct {
	state  byte
	passed bool
	r, c   int
}

func (m *Maze) clone() [][]Node {
	newNodes := make([][]Node, 0, len(m.node))
	for i, _ := range m.node {
		newLine := make([]Node, len(m.node[i]))
		copy(newLine, m.node[i])
		newNodes = append(newNodes, newLine)
	}
	return newNodes
}

func (m *Maze) mimicAnswer() [][]Node {
	mazeNodes := m.clone()
	for _, n := range m.answers {
		mazeNodes[n.r][n.c].state = []byte("X")[0]
	}
	return mazeNodes
}

func (m *Maze) dfs() bool {
	// find start point
	var rStart, cStart int
	for r, nodeLine := range m.node {
		for c, nodeOne := range nodeLine {
			if nodeOne.state == []byte("S")[0] {
				rStart, cStart = r, c
			}
		}
	}

	if found, _ := m.internalDfs(rStart, cStart); found {
		return true
	} else {
		return false
	}
}

func (m *Maze) internalDfs(r, c int) (bool, *Node) {
	if r < 0 || c < 0 || r >= len(m.node) || c >= len(m.node[r]) {
		return false, nil
	}
	if m.node[r][c].passed || m.node[r][c].state == []byte("|")[0] || m.node[r][c].state == []byte("-")[0] {
		return false, nil
	}
	m.node[r][c].passed = true
	if m.node[r][c].state == []byte("E")[0] {
		return true, &m.node[r][c]
	}
	var found bool
	var next *Node
	if found, next = m.internalDfs(r-1, c); found {
		m.answers = append(m.answers, next)
	} else if found, next = m.internalDfs(r, c-1); found {
		m.answers = append(m.answers, next)
	} else if found, next = m.internalDfs(r+1, c); found {
		m.answers = append(m.answers, next)
	} else if found, next = m.internalDfs(r, c+1); found {
		m.answers = append(m.answers, next)
	}
	return found, &m.node[r][c]
}

func (m *Maze) create(in []byte) {
	tmpRow := make([]Node, 0)
	countR, countC := 0, 0
	for _, c := range in {
		if c == []byte("\n")[0] {
			m.node = append(m.node, tmpRow)
			tmpRow = make([]Node, 0)
			countR++
			countC = 0
		} else {
			tmpRow = append(tmpRow, Node{state: c, passed: false, r: countR, c: countC})
			countC++
		}
	}
}

func (m *Maze) printNodes(nodes [][]Node) {
	for _, nodeRow := range nodes {
		for _, nodeOne := range nodeRow {
			fmt.Printf("%c", nodeOne.state)
		}
		fmt.Println()
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	maze := Maze{}
	maze.create(bytes)
	fmt.Println("[\t\tinput\t\t ]")
	maze.printNodes(maze.node)
	// solution
	fmt.Println("[\t\tOutput\t\t ]")
	if maze.dfs() {
		// output
		ansNodes := maze.mimicAnswer()
		maze.printNodes(ansNodes)
	} else {
		fmt.Println("No answers")
	}
}
