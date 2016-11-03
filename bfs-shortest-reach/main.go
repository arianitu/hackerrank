package main

import "fmt"

type Node struct {
	neighbours []*Node
	data       int
	distance   int
}

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

func consumeQuery() {
	nodes := make(map[int]*Node)

	var numberOfNodes, edges int
	fmt.Scan(&numberOfNodes, &edges)

	for i := 0; i < numberOfNodes; i++ {
		nodes[i+1] = &Node{data: i + 1}
	}

	for i := 0; i < edges; i++ {
		var from, to int
		fmt.Scan(&from, &to)
		nodes[from].neighbours = append(nodes[from].neighbours, nodes[to])
		nodes[to].neighbours = append(nodes[to].neighbours, nodes[from])
	}
	var startingNode int
	fmt.Scan(&startingNode)

	start := nodes[startingNode]
	q := NewQueue()
	visited := make(map[int]bool)

	q.Push(start)
	for !q.IsEmpty() {
		node := q.Pop()

		for _, neighbour := range node.neighbours {
			if visited[neighbour.data] {
				continue
			}
			neighbour.distance = node.distance + 6
			q.Push(neighbour)

			visited[neighbour.data] = true
		}
	}
	for i := 0; i < numberOfNodes; i++ {
		node := nodes[i+1]
		if node.data == startingNode {
			continue
		}
		if node.distance <= 0 {
			fmt.Print("-1 ")
		} else {
			fmt.Print(node.distance, " ")
		}
	}
	fmt.Print("\n")

}

func main() {
	var q int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		consumeQuery()
	}
}
