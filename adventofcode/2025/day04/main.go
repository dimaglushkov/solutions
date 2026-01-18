package main

import (
	"fmt"

	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func process(input []string) [][]bool {
	n, m := len(input), len(input[0])

	data := make([][]bool, n)
	for i := range data {
		data[i] = make([]bool, m)

		for j := range data[i] {
			data[i][j] = input[i][j] == '@'
		}
	}

	return data
}

func part1(input []string) {
	var res int

	matrix := process(input)
	n, m := len(matrix), len(matrix[0])

	for i := range matrix {
		for j := range matrix[i] {
			if !matrix[i][j] {
				continue
			}

			cnt := 0
			dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

			for _, d := range dirs {
				if dx, dy := i+d[0], j+d[1]; dx >= 0 && dx < n && dy >= 0 && dy < m && matrix[dx][dy] {
					cnt++
				}
			}

			if cnt < 4 {
				res++
			}
		}
	}

	fmt.Println(res)
}

func process2(input []string) [][]int {
	n, m := len(input), len(input[0])

	data := make([][]int, n)
	for i := range data {
		data[i] = make([]int, m)

		for j := range data[i] {
			if input[i][j] == '@' {
				data[i][j] = 1
			}
		}
	}

	return data
}

func part2(input []string) {
	var res int

	matrix := process2(input)
	n, m := len(matrix), len(matrix[0])

	changed := true
	for changed {
		changed = false

		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] != 1 {
					continue
				}

				cnt := 0
				dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

				for _, d := range dirs {
					if dx, dy := i+d[0], j+d[1]; dx >= 0 && dx < n && dy >= 0 && dy < m && (matrix[dx][dy] == 1 || matrix[dx][dy] == -1) {
						cnt++
					}
				}

				if cnt < 4 {
					changed = true
					matrix[i][j] = -1
				}
			}
		}

		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] == -1 {
					matrix[i][j] = 0
					res++
				}
			}
		}
	}

	fmt.Println(res)
}

func main() {
	exampleLines := util.ReadInput(true)

	fmt.Println("Example:")
	part1(exampleLines)
	part2(exampleLines)

	lines := util.ReadInput(false)
	fmt.Println("=========")
	part1(lines)
	part2(lines)
}
