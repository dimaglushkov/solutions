package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
)

var dirs = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

type pair struct {
	x, y int
}

type queue []pair

func (q *queue) push(x pair) {
	*q = append(*q, x)
}
func (q *queue) pop() pair {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func part1(matrix [][]int) {
	var res int

	trailheads := make(map[pair]map[pair]bool)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				head := pair{i, j}

				trailheads[head] = make(map[pair]bool)

				q := queue{}
				q.push(head)

				for len(q) > 0 {
					p := q.pop()

					if matrix[p.x][p.y] == 9 {
						trailheads[head][p] = true
					}

					for _, d := range dirs {
						if dx, dy := p.x+d[0], p.y+d[1]; dx >= 0 && dx < len(matrix) && dy >= 0 && dy < len(matrix[0]) && matrix[dx][dy] == matrix[p.x][p.y]+1 {
							q.push(pair{dx, dy})
						}
					}

				}
			}
		}
	}

	for _, v := range trailheads {
		res += len(v)
	}

	fmt.Println(res)
}

func part2(matrix [][]int) {
	var res int

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				head := pair{i, j}

				q := queue{}
				q.push(head)

				for len(q) > 0 {
					p := q.pop()

					if matrix[p.x][p.y] == 9 {
						res++
						continue
					}

					for _, d := range dirs {
						if dx, dy := p.x+d[0], p.y+d[1]; dx >= 0 && dx < len(matrix) && dy >= 0 && dy < len(matrix[0]) && matrix[dx][dy] == matrix[p.x][p.y]+1 {
							q.push(pair{dx, dy})
						}
					}

				}
			}
		}
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	matrix := make([][]int, len(lines))
	for i, l := range lines {
		matrix[i] = make([]int, len(l))
		for j, c := range l {
			matrix[i][j] = int(c - '0')
		}
	}

	part1(matrix)
	part2(matrix)
}
