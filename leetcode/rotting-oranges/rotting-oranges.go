package main

import "fmt"

// source: https://leetcode.com/problems/rotting-oranges/

type orange struct {
	x, y int
	dist int
}

func orangesRotting(grid [][]int) int {
	q := make([]orange, 0)
	push := func(x orange) {
		q = append(q, x)
	}
	pop := func() orange {
		x := q[0]
		q = q[1:]
		return x
	}

	cnt := 0
	res := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				cnt++
			}
			if grid[i][j] == 2 {
				push(orange{i, j, 0})
			}
		}
	}

	for len(q) > 0 {
		data := pop()
		x, y, dist := data.x, data.y, data.dist
		if dist > res {
			res = dist
		}

		grid[x][y] = 2
		dist++
		if dx := x - 1; dx >= 0 && grid[dx][y] == 1 {
			cnt--
			push(orange{dx, y, dist})
			grid[dx][y] = 3
		}
		if dx := x + 1; dx < n && grid[dx][y] == 1 {
			cnt--
			push(orange{dx, y, dist})
			grid[dx][y] = 3
		}
		if dy := y - 1; dy >= 0 && grid[x][dy] == 1 {
			cnt--
			push(orange{x, dy, dist})
			grid[x][dy] = 3
		}
		if dy := y + 1; dy < m && grid[x][dy] == 1 {
			cnt--
			push(orange{x, dy, dist})
			grid[x][dy] = 3
		}
	}

	if cnt == 0 {
		return res
	}
	return -1
}

func main() {
	var grid = [][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}
	fmt.Println("Expected: 1	Output: ", orangesRotting(grid))

	var grid0 = [][]int{{2, 1, 1}, {1, 1, 1}, {0, 1, 2}}
	fmt.Println("Expected: 2	Output: ", orangesRotting(grid0))

	// Example 1
	var grid1 = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println("Expected: 4	Output: ", orangesRotting(grid1))

	// Example 2
	var grid2 = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Println("Expected: -1	Output: ", orangesRotting(grid2))

	// Example 3
	var grid3 = [][]int{{0, 2}}
	fmt.Println("Expected: 0	Output: ", orangesRotting(grid3))

}
