package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/01-matrix/

type queue struct {
	val [][2]int
}

func newQueue() queue {
	return queue{
		val: make([][2]int, 0),
	}
}

func (q *queue) push(i, j int) {
	q.val = append(q.val, [2]int{i, j})
}

func (q *queue) pop() (int, int) {
	i, j := q.val[0][0], q.val[0][1]
	q.val = q.val[1:]
	return i, j
}

func (q *queue) len() int {
	return len(q.val)
}

func updateMatrix(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	mem := make([][]int, n)
	q := newQueue()

	for i := range mem {
		mem[i] = make([]int, m)
		for j := range mem[i] {
			if mat[i][j] == 0 {
				q.push(i, j)
			} else {
				mem[i][j] = math.MaxInt
			}
		}
	}

	for q.len() > 0 {
		i, j := q.pop()
		v := mem[i][j] + 1

		if i-1 >= 0 && mem[i-1][j] > v {
			mem[i-1][j] = v
			q.push(i-1, j)
		}
		if i+1 < n && mem[i+1][j] > v {
			mem[i+1][j] = v
			q.push(i+1, j)
		}

		if j-1 >= 0 && mem[i][j-1] > v {
			mem[i][j-1] = v
			q.push(i, j-1)
		}
		if j+1 < m && mem[i][j+1] > v {
			mem[i][j+1] = v
			q.push(i, j+1)
		}

	}

	return mem
}

func main() {
	// Example 1
	var mat1 = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println("Expected: [[0,0,0],[0,1,0],[0,0,0]]	Output: ", updateMatrix(mat1))

	// Example 2
	var mat2 = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println("Expected: [[0,0,0],[0,1,0],[1,2,1]]	Output: ", updateMatrix(mat2))

}
