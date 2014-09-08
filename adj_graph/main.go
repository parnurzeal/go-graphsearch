package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Maze byte

func (m Maze) Compare() State {
	switch m {
	case 'S':
		return START
	case 'E':
		return END
	case ' ':
		return SPACE
	default:
		return OBSTACLE
	}
}

func createMazeGraph(in []byte) Graph {
	g := Graph{nodes: make(map[int]Node)}
	// create matrix for easy translating to graph
	matrix := [][]Maze{}
	line := []Maze{}
	for _, char := range in {
		if char != '\n' {
			line = append(line, Maze(char))
		} else {
			matrix = append(matrix, line)
			line = []Maze{}
		}
	}
	// create graph from matrix
	count := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			g.addNode(count, matrix[r][c])
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
	return g
}

func mimicAnswers(in []byte, answers []int) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	count := 0
	for idx, char := range in {
		for _, tmp := range answers {
			if tmp == count {
				out[idx] = 'X'
			}
		}
		if char != '\n' {
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
	g := createMazeGraph(inBytes)
	bfsAns := g.bfs()
	fmt.Println(string(mimicAnswers(inBytes, bfsAns)))
	dfsAns := g.dfs()
	fmt.Println(string(mimicAnswers(inBytes, dfsAns)))
}
