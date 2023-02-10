package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/as-far-from-land-as-possible/

type xyd struct {
	x, y, d int
}

type queue []xyd

func (q *queue) push(x xyd) {
	*q = append(*q, x)
}
func (q *queue) pop() xyd {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDistance(grid [][]int) int {
	n := len(grid)
	visited := make([][]bool, n)
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	q := new(queue)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				q.push(xyd{i, j, 0})
			}
		}
	}
	if len(*q) == n*n {
		return -1
	}

	ans := -1
	for len(*q) > 0 {
		p := q.pop()
		if visited[p.x][p.y] {
			continue
		}
		visited[p.x][p.y] = true
		ans = max(ans, p.d)
		for _, dir := range dirs {
			if nx, ny := p.x+dir[0], p.y+dir[1]; nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				q.push(xyd{nx, ny, p.d + 1})
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
			want: 2,
		},
		{
			grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want: 4,
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := maxDistance(tc.grid)
		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
