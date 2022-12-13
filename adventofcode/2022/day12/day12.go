package main

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"strings"
)

type point struct {
	x, y, s int
}

type queue []point

func (q *queue) push(x, y, s int) {
	*q = append(*q, point{x, y, s})
}
func (q *queue) pop() (int, int, int) {
	x := (*q)[0]
	*q = (*q)[1:]
	return x.x, x.y, x.s
}

func dayTwelve1() {
	text := util.ReadInput("day12.input")
	var sx, sy, ex, ey int
	n, m := len(text), len(text[0])
	steps := make([][]int, n)
	for i := range text {
		steps[i] = make([]int, m)
		for j := range text[i] {
			steps[i][j] = 1 << 32

			if text[i][j] == 'S' {
				sx, sy = i, j
			}
			if text[i][j] == 'E' {
				ex, ey = i, j
			}
		}
		if strings.ContainsRune(text[i], 'S') {
			text[i] = strings.Replace(text[i], "S", "a", 1)
		}
		if strings.ContainsRune(text[i], 'E') {
			text[i] = strings.Replace(text[i], "E", "z", 1)
		}
	}

	q := make(queue, 0)
	q.push(sx, sy, 0)
	for len(q) > 0 {
		x, y, s := q.pop()
		if s >= steps[x][y] {
			continue
		}
		steps[x][y] = s
		if x == ex && y == ey {
			continue
		}

		s++
		// up
		if nx := x - 1; nx >= 0 && text[x][y]+1 >= text[nx][y] && steps[nx][y] > s {
			q.push(nx, y, s)
		}
		// down
		if nx := x + 1; nx < n && text[x][y]+1 >= text[nx][y] && steps[nx][y] > s {
			q.push(nx, y, s)
		}
		// left
		if ny := y - 1; ny >= 0 && text[x][y]+1 >= text[x][ny] && steps[x][ny] > s {
			q.push(x, ny, s)
		}
		// right
		if ny := y + 1; ny < m && text[x][y]+1 >= text[x][ny] && steps[x][ny] > s {
			q.push(x, ny, s)
		}
	}

	println(steps[ex][ey])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dayTwelve2() {
	text := util.ReadInput("day12.input")
	var sx, sy int
	n, m := len(text), len(text[0])
	steps := make([][]int, n)
	for i := range text {
		steps[i] = make([]int, m)
		for j := range text[i] {
			steps[i][j] = 1 << 32

			if text[i][j] == 'E' {
				sx, sy = i, j
			}
		}
		if strings.ContainsRune(text[i], 'S') {
			text[i] = strings.Replace(text[i], "S", "a", 1)
		}
		if strings.ContainsRune(text[i], 'E') {
			text[i] = strings.Replace(text[i], "E", "z", 1)
		}
	}

	q := make(queue, 0)
	q.push(sx, sy, 0)
	res := 1 << 32
	for len(q) > 0 {
		x, y, s := q.pop()
		if s >= steps[x][y] {
			continue
		}
		steps[x][y] = s
		if text[x][y] == 'a' {
			res = min(res, s)
			continue
		}

		s++
		// up
		if nx := x - 1; nx >= 0 && text[x][y]-1 <= text[nx][y] && steps[nx][y] > s {
			q.push(nx, y, s)
		}
		// down
		if nx := x + 1; nx < n && text[x][y]-1 <= text[nx][y] && steps[nx][y] > s {
			q.push(nx, y, s)
		}
		// left
		if ny := y - 1; ny >= 0 && text[x][y]-1 <= text[x][ny] && steps[x][ny] > s {
			q.push(x, ny, s)
		}
		// right
		if ny := y + 1; ny < m && text[x][y]-1 <= text[x][ny] && steps[x][ny] > s {
			q.push(x, ny, s)
		}
	}

	println(res)
}

func main() {
	dayTwelve1()
	dayTwelve2()
}
