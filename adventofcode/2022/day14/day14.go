package main

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"strconv"
	"strings"
)

const (
	rock         = '#'
	air          = '.'
	chillingSand = 'o'
	startingSand = '+'
)

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dayFourteen1() {
	// parse input
	text := util.ReadInput("day14.input")
	n, m := -1, -1
	input := make([][][2]int, len(text))
	for i, l := range text {
		coords := strings.Split(l, " -> ")
		input[i] = make([][2]int, len(coords))
		for j := range coords {
			xy := strings.Split(coords[j], ",")
			x, y := atoi(xy[0]), atoi(xy[1])
			input[i][j] = [2]int{x, y}
			n, m = max(n, y), max(m, x)
		}
	}
	n++
	m++

	// build matrix
	matrix := make([][]byte, n+1)
	for i := range matrix {
		matrix[i] = make([]byte, m+1)
		for j := range matrix[i] {
			matrix[i][j] = air
		}
	}
	for i := range input {
		var s, e [2]int
		s = input[i][0]
		for j := 1; j < len(input[i]); j++ {
			e = input[i][j]
			if s[0] == e[0] {
				d := 1
				if s[1] > e[1] {
					d = -1
				}
				for r := s[1]; r != e[1]; r += d {
					matrix[r][s[0]] = rock
				}
				matrix[e[1]][s[0]] = rock
			} else {
				d := 1
				if s[0] > e[0] {
					d = -1
				}
				for c := s[0]; c != e[0]; c += d {
					matrix[s[1]][c] = rock
				}
				matrix[s[1]][e[0]] = rock
			}
			s = e
		}
	}

	var drip func() (bool, int, int)
	drip = func() (bool, int, int) {
		row, col := 0, 500

		for row < n && col < m && col >= 0 {
			if matrix[row+1][col] == air {
				row++
			} else if matrix[row+1][col-1] == air {
				col--
			} else if matrix[row+1][col+1] == air {
				col++
			} else {
				break
			}
		}
		return row < n && col < m && col >= 0, row, col
	}

	res := 0
	for s, r, c := drip(); s; s, r, c = drip() {
		matrix[r][c] = chillingSand
		res++
	}

	//for i := range matrix {
	//	for j := 494; j < m; j++ {
	//		fmt.Printf("%c", matrix[i][j])
	//	}
	//	println()
	//}

	println(res)
}

func dayFourteen2() {
	// parse input
	text := util.ReadInput("day14.input")
	n, m := -1, -1
	input := make([][][2]int, len(text))
	for i, l := range text {
		coords := strings.Split(l, " -> ")
		input[i] = make([][2]int, len(coords))
		for j := range coords {
			xy := strings.Split(coords[j], ",")
			x, y := atoi(xy[0]), atoi(xy[1])
			input[i][j] = [2]int{x, y}
			n, m = max(n, y), max(m, x)
		}
	}
	n += 3
	m += n + 1

	// build matrix
	matrix := make([][]byte, n)
	for i := range matrix {
		matrix[i] = make([]byte, m)
		for j := range matrix[i] {
			if i == n-1 {
				matrix[i][j] = rock
			} else {
				matrix[i][j] = air
			}
		}
	}
	for i := range input {
		var s, e [2]int
		s = input[i][0]
		for j := 1; j < len(input[i]); j++ {
			e = input[i][j]
			if s[0] == e[0] {
				d := 1
				if s[1] > e[1] {
					d = -1
				}
				for r := s[1]; r != e[1]; r += d {
					matrix[r][s[0]] = rock
				}
				matrix[e[1]][s[0]] = rock
			} else {
				d := 1
				if s[0] > e[0] {
					d = -1
				}
				for c := s[0]; c != e[0]; c += d {
					matrix[s[1]][c] = rock
				}
				matrix[s[1]][e[0]] = rock
			}
			s = e
		}
	}

	var drip func() (int, int)
	drip = func() (int, int) {
		prow, pcol := -1, -1
		row, col := 0, 500

		for row != prow || col != pcol {
			prow, pcol = row, col
			if row+1 < n && matrix[row+1][col] == air {
				row++
			} else if row+1 < n && col-1 >= 0 && matrix[row+1][col-1] == air {
				row++
				col--
			} else if row+1 < n && col+1 < m && matrix[row+1][col+1] == air {
				row++
				col++
			}
		}

		return row, col
	}

	res := 0
	for matrix[0][500] != chillingSand {
		r, c := drip()
		matrix[r][c] = chillingSand
		res++
	}

	//for i := range matrix {
	//	for j := 485; j < m; j++ {
	//		fmt.Printf("%c", matrix[i][j])
	//	}
	//	println()
	//}

	println(res)
}

func main() {
	dayFourteen1()
	dayFourteen2()
}
