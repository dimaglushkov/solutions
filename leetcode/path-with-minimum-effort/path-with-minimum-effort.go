package main

import (
	"container/heap"
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/path-with-minimum-effort/

type Heap [][3]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumEffortPath(heights [][]int) int {
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	h, w := len(heights), len(heights[0])

	minEfforts := make([][]int, h)
	for i := 0; i < h; i++ {
		minEfforts[i] = make([]int, w)
		for j := 0; j < w; j++ {
			minEfforts[i][j] = math.MaxInt
		}
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	pq := &Heap{{0, 0, 0}}
	heap.Init(pq)

	for {
		edge := heap.Pop(pq).([3]int)
		if edge[0] > minEfforts[edge[1]][edge[2]] {
			continue
		}

		if edge[1] == h-1 && edge[2] == w-1 {
			return edge[0]
		}

		for _, dir := range directions {
			di, dj := edge[1]+dir[0], edge[2]+dir[1]
			if di >= 0 && di < h && dj >= 0 && dj < w {
				effort := max(edge[0], abs(heights[di][dj]-heights[edge[1]][edge[2]]))
				if effort < minEfforts[di][dj] {
					minEfforts[di][dj] = effort
					heap.Push(pq, [3]int{effort, di, dj})
				}
			}
		}

	}
}

func minimumEffortPath_(heights [][]int) int {
	h, w := len(heights), len(heights[0])

	minEfforts := make([][]int, h)
	for i := 0; i < h; i++ {
		minEfforts[i] = make([]int, w)
		for j := 0; j < w; j++ {
			minEfforts[i][j] = math.MaxInt
		}
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var util func(i, j, prev, effort int)
	util = func(i, j, prev, effort int) {
		if i < 0 || j < 0 || i >= h || j >= w || effort >= minEfforts[i][j] {
			return
		}

		if newEffort := abs(prev - heights[i][j]); newEffort > effort {
			effort = newEffort
		}
		if effort < minEfforts[i][j] {
			minEfforts[i][j] = effort
		}

		util(i+1, j, heights[i][j], effort)
		util(i, j+1, heights[i][j], effort)
		util(i-1, j, heights[i][j], effort)
		util(i, j-1, heights[i][j], effort)
	}

	util(0, 0, heights[0][0], 0)
	return minEfforts[h-1][w-1]
}

func main() {
	// Example 1
	var heights1 = [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println("Expected: 2	Output: ", minimumEffortPath(heights1))

	// Example 4
	var heights4 = [][]int{
		{8, 3, 2, 5, 2, 10, 7, 1, 8, 9},
		{1, 4, 9, 1, 10, 2, 4, 10, 3, 5},
		{4, 10, 10, 3, 6, 1, 3, 9, 8, 8},
		{4, 4, 6, 10, 10, 10, 2, 10, 8, 8},
		{9, 10, 2, 4, 1, 2, 2, 6, 5, 7},
		{2, 9, 2, 6, 1, 4, 7, 6, 10, 9},
		{8, 8, 2, 10, 8, 2, 3, 9, 5, 3},
		{2, 10, 9, 3, 5, 1, 7, 4, 5, 6},
		{2, 3, 9, 2, 5, 10, 2, 7, 1, 8},
		{9, 10, 4, 10, 7, 4, 9, 3, 1, 6},
	}
	fmt.Println("Expected: 5	Output: ", minimumEffortPath(heights4))

	// Example 2
	var heights2 = [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	fmt.Println("Expected: 1	Output: ", minimumEffortPath(heights2))

	// Example 3
	var heights3 = [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	fmt.Println("Expected: 0	Output: ", minimumEffortPath(heights3))

}
