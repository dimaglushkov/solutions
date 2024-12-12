package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
	"strconv"
)

var dirs = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func part1(matrix [][]byte) {
	var res int
	n, m := len(matrix), len(matrix[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}

	area := make(map[string]int)
	perimeter := make(map[string]int)

	var explore func(i, j int, id string)
	explore = func(i, j int, id string) {
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		x := matrix[i][j]
		area[id]++

		for _, d := range dirs {
			if di, dj := i+d[0], j+d[1]; di >= 0 && di < n && dj >= 0 && dj < m && matrix[di][dj] == x {
				explore(di, dj, id)
			} else {
				perimeter[id]++
			}
		}
	}

	for i := range n {
		for j := range m {
			if !visited[i][j] {
				explore(i, j, string(matrix[i][j])+strconv.Itoa(i)+"+"+strconv.Itoa(j))
			}
		}
	}

	for i := range area {
		res += area[i] * perimeter[i]
	}

	fmt.Println(res)
}

func part2(matrix [][]byte) {
	var res int
	n, m := len(matrix), len(matrix[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}

	area := make(map[string]int)
	corners := make(map[string]int)

	countCorners := func(i, j int) int {
		x := matrix[i][j]
		cntCorner := 0

		// outer top left
		if di, dj := i-1, j-1; (di < 0 || matrix[di][j] != x) && (dj < 0 || matrix[i][dj] != x) {
			cntCorner++
		}
		// outer top right
		if di, dj := i-1, j+1; (di < 0 || matrix[di][j] != x) && (dj >= m || matrix[i][dj] != x) {
			cntCorner++
		}
		// outer bot left
		if di, dj := i+1, j-1; (di >= n || matrix[di][j] != x) && (dj < 0 || matrix[i][dj] != x) {
			cntCorner++
		}
		// outer bot right
		if di, dj := i+1, j+1; (di >= n || matrix[di][j] != x) && (dj >= m || matrix[i][dj] != x) {
			cntCorner++
		}

		// inner
		for _, d := range [][2]int{{1, -1}, {1, 1}, {-1, 1}, {-1, -1}} {
			if di, dj := i+d[0], j+d[1]; di >= 0 && di < n && dj >= 0 && dj < m && matrix[i][dj] == x && matrix[di][j] == x && matrix[di][dj] != x {
				cntCorner++
			}
		}

		return cntCorner
	}

	var explore func(i, j int, id string)
	explore = func(i, j int, id string) {
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		x := matrix[i][j]
		area[id]++

		corners[id] += countCorners(i, j)

		for _, d := range dirs {
			if di, dj := i+d[0], j+d[1]; di >= 0 && di < n && dj >= 0 && dj < m && matrix[di][dj] == x {
				explore(di, dj, id)
			}
		}
	}

	for i := range n {
		for j := range m {
			if !visited[i][j] {
				explore(i, j, string(matrix[i][j])+strconv.Itoa(i)+"+"+strconv.Itoa(j))
			}
		}
	}

	for i := range area {
		res += area[i] * corners[i]
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	matrix := make([][]byte, len(lines))
	for i := range matrix {
		matrix[i] = make([]byte, len(lines[i]))
		for j := range matrix[i] {
			matrix[i][j] = lines[i][j]
		}
	}

	part1(matrix)
	part2(matrix)
}
