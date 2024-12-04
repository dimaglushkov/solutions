package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func part1(input [][]byte) {
	var res int

	const xmas = "XMAS"

	var dirs = [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	var check func(x, y int, dir [2]int, depth int) bool
	check = func(x, y int, dir [2]int, depth int) bool {
		if depth == 4 {
			return true
		}

		if dx, dy := x+dir[0], y+dir[1]; dx >= 0 && dx < len(input) && dy >= 0 && dy < len(input[0]) && input[dx][dy] == xmas[depth] {
			return check(dx, dy, dir, depth+1)
		}

		return false
	}

	for i := range input {
		for j := range input[i] {
			if input[i][j] != 'X' {
				continue
			}

			for _, d := range dirs {
				if check(i, j, d, 1) {
					res++
				}
			}
		}
	}

	fmt.Println(res)
}

func part2(input [][]byte) {
	var res int

	cnt := make(map[[2]int]int)

	const mas = "MAS"
	var dirs = [][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	var check func(x, y int, dir [2]int, depth int) bool
	check = func(x, y int, dir [2]int, depth int) bool {
		if depth == 3 {
			return true
		}

		if dx, dy := x+dir[0], y+dir[1]; dx >= 0 && dx < len(input) && dy >= 0 && dy < len(input[0]) && input[dx][dy] == mas[depth] {
			return check(dx, dy, dir, depth+1)
		}

		return false
	}

	for i := range input {
		for j := range input[i] {
			if input[i][j] != 'M' {
				continue
			}

			for _, d := range dirs {
				if check(i, j, d, 1) {
					cnt[[2]int{i + d[0], j + d[1]}]++
				}
			}
		}
	}

	for _, c := range cnt {
		if c > 1 {
			res++
		}
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	var matrix [][]byte
	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}

	part1(matrix)
	part2(matrix)
}
