package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
)

const (
	up = iota
	right
	down
	left
	space
	blocker
	guard
	checked
)

var dir = [][2]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}} // left

func part1(x, y int, matrix [][]int) [][]bool {
	var res int
	var cd = 0
	var visited = util.MakeMatrix[bool](len(matrix), len(matrix[0]))

	move := func() bool {
		matrix[x][y] = checked
		visited[x][y] = true
		dx, dy := x+dir[cd][0], y+dir[cd][1]
		if dx < 0 || dx >= len(matrix) || dy < 0 || dy >= len(matrix[0]) {
			return false
		}

		if matrix[dx][dy] == blocker {
			cd++
			if cd == len(dir) {
				cd = 0
			}
			return true
		}

		x, y = dx, dy

		return true
	}

	for move() {
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == checked {
				res++
			}
		}
	}

	fmt.Println(res)

	return visited
}

func part2(x, y int, matrix [][]int, visited [][]bool) {
	var res int
	isCycling := func(x, y, d int, c [][]int) bool {
		loop := false
		var routes [4]map[[2]int]bool

		for i := range dir {
			routes[i] = make(map[[2]int]bool)
		}

		move := func() bool {
			dx, dy := x+dir[d][0], y+dir[d][1]
			if dx < 0 || dx >= len(c) || dy < 0 || dy >= len(c[0]) {
				return false
			}

			if routes[d][[2]int{x, y}] {
				loop = true
				return false
			}

			routes[d][[2]int{x, y}] = true

			if c[dx][dy] == blocker {
				d++
				if d == len(dir) {
					d = 0
				}
				return true
			}

			x, y = dx, dy

			return true
		}

		for move() {
		}

		return loop
	}

	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				matrix[i][j] = blocker
				if isCycling(x, y, 0, matrix) {
					res++
				}
				matrix[i][j] = space
			}
		}
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	var x, y int
	matrix := make([][]int, len(lines))
	for i, l := range lines {
		matrix[i] = make([]int, len(l))
		for j := range l {
			switch l[j] {
			case '.':
				matrix[i][j] = space
			case '#':
				matrix[i][j] = blocker
			case '^':
				matrix[i][j] = guard
				x, y = i, j
			}
		}
	}

	visited := part1(x, y, matrix)
	part2(x, y, matrix, visited)
}

// doesn't work unfortunately
func part2Attempt1(x, y int, matrix [][]int) {
	var res int
	var cd = 0

	var routes [4]map[[2]int]bool

	for d := range dir {
		routes[d] = make(map[[2]int]bool)
	}

	nextDir := func() int {
		if cd+1 == len(dir) {
			return 0
		}
		return cd + 1
	}

	check := func() bool {
		nd := nextDir()
		for i, j := x, y; i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]); i, j = i+dir[nd][0], j+dir[nd][1] {
			if routes[nd][[2]int{i, j}] {
				return true
			}
		}
		return false
	}

	move := func() bool {
		routes[cd][[2]int{x, y}] = true
		if check() {
			res++
		}

		dx, dy := x+dir[cd][0], y+dir[cd][1]
		if dx < 0 || dx >= len(matrix) || dy < 0 || dy >= len(matrix[0]) {
			return false
		}

		matrix[x][y] = checked

		if matrix[dx][dy] == blocker {
			cd = nextDir()
			return true
		}
		x, y = dx, dy

		return true
	}

	for move() {
	}

	fmt.Println(res)
}
