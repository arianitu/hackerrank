package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	incrementingCounter = 0
)

type Queue struct {
	nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{nodes: make([]*Node, 0)}
}
func (q *Queue) Push(node *Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) Pop() *Node {
	r := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	return r
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

type Node struct {
	id    int
	data  int
	edges []*Node
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func consumeRow(scanner *bufio.Scanner) []*Node {
	nodes := make([]*Node, 0)

	ok := scanner.Scan()
	if !ok {
		panic("failed to scan")
	}

	columnsAsString := strings.Split(scanner.Text(), " ")
	for _, columnString := range columnsAsString {
		columnInt, err := strconv.Atoi(columnString)
		checkError(err)

		nodes = append(nodes, &Node{id: incrementingCounter, data: columnInt})
		incrementingCounter += 1
	}

	return nodes
}

var directions = [][]int{
	// left
	{-1, 0},
	// right
	{1, 0},
	// top
	{0, -1},
	// bottom
	{0, 1},
	// top left
	{-1, 1},
	// top right
	{1, 1},
	// bottom left
	{-1, -1},
	// bottom right
	{1, -1},
}

func addEdges(matrix [][]*Node, row int, column int, node *Node) {
	height := len(matrix) - 1
	width := len(matrix[0]) - 1

	for _, direction := range directions {
		if column+direction[0] >= 0 && column+direction[0] <= width &&
			row+direction[1] >= 0 && row+direction[1] <= height {

			to := matrix[row+direction[1]][column+direction[0]]

			node.edges = append(node.edges, to)
			to.edges = append(to.edges, node)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	unvisited := make(map[int]*Node)

	var rows, columns int
	fmt.Scanln(&rows)
	fmt.Scanln(&columns)

	matrix := make([][]*Node, rows)
	for r := 0; r < rows; r++ {
		row := consumeRow(scanner)
		matrix[r] = row
	}

	for r := 0; r < rows; r++ {
		row := matrix[r]
		for c, node := range row {
			if node.data == 1 {
				unvisited[node.id] = node
			}
			addEdges(matrix, r, c, node)
		}
	}

	mostConnected := 0
	q := NewQueue()
	for _, nodeToVisit := range unvisited {
		distance := 1
		q.Push(nodeToVisit)
		delete(unvisited, nodeToVisit.id)

		for !q.IsEmpty() {
			node := q.Pop()
			for _, edge := range node.edges {
				if unvisited[edge.id] == nil {
					continue

				}

				q.Push(edge)
				distance += 1
				delete(unvisited, edge.id)
			}
		}
		if distance > mostConnected {
			mostConnected = distance
		}
	}

	fmt.Println(mostConnected)
}
