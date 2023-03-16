package main

// source: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type pair struct {
	x *Node
	y int
}

type queue []pair

func (q *queue) push(x *Node, y int) {
	*q = append(*q, pair{x: x, y: y})
}
func (q *queue) pop() (*Node, int) {
	x := (*q)[0]
	*q = (*q)[1:]
	return x.x, x.y
}
func (q *queue) peek() (*Node, int) {
	x := (*q)[0]
	return x.x, x.y
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var q queue
	q.push(root, 0)
	for len(q) > 0 {
		x, d := q.pop()
		if x == nil {
			break
		}

		q.push(x.Left, d+1)
		q.push(x.Right, d+1)

		if n, nd := q.peek(); nd == d {
			x.Next = n
		}
	}
	return root
}
