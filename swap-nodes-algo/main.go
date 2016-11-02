package main

import "fmt"

type Queue struct {
	nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{nodes: make([]*Node, 0)}
}

func (q *Queue) Push(node *Node) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

func (q *Queue) Pop() *Node {
	returnValue := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]

	return returnValue
}

type Node struct {
	data  int
	depth int
	left  *Node
	right *Node
}

func swapChildren(node *Node) {
	tmp := node.left
	node.left = node.right
	node.right = tmp
}

func consumeNodes() (*Node, *Node) {
	var leftData int
	var rightData int
	var left *Node = nil
	var right *Node = nil

	fmt.Scan(&leftData, &rightData)
	if leftData != -1 {
		left = &Node{data: leftData}
	}
	if rightData != -1 {
		right = &Node{data: rightData}
	}

	return left, right

}

func InOrder(node *Node) {
	if node == nil {
		return
	}

	InOrder(node.left)
	fmt.Printf("%v ", node.data)
	InOrder(node.right)
}

func shouldSwapAtDepth(K int, depth int) bool {
	if depth%K == 0 {
		return true
	}
	return false
}

func main() {
	root := &Node{data: 1, depth: 1}
	numberOfNodes := 0
	fmt.Scan(&numberOfNodes)

	q := NewQueue()
	q.Push(root)
	for i := 0; i < numberOfNodes; i++ {
		n := q.Pop()

		left, right := consumeNodes()
		n.left = left
		n.right = right

		if n.left != nil {
			q.Push(n.left)
		}
		if n.right != nil {
			q.Push(n.right)
		}
	}

	var numberOfSwaps int
	fmt.Scan(&numberOfSwaps)
	for t := 0; t < numberOfSwaps; t++ {
		var K int
		fmt.Scan(&K)

		q := NewQueue()
		root.depth = 1
		q.Push(root)

		for !q.IsEmpty() {
			n := q.Pop()
			if shouldSwapAtDepth(K, n.depth) {
				swapChildren(n)
			}
			if n.left != nil {
				n.left.depth = n.depth + 1
				q.Push(n.left)
			}
			if n.right != nil {
				n.right.depth = n.depth + 1
				q.Push(n.right)
			}
		}

		InOrder(root)
		fmt.Print("\n")
	}
}
