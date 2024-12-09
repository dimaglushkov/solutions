package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func check(p1, p2 [2]int, n, m int) [][2]int {
	var ans [][2]int
	dx, dy := p1[0]-p2[0], p1[1]-p2[1]

	if c1x, c1y := p1[0]+dx, p1[1]+dy; c1x >= 0 && c1x < n && c1y >= 0 && c1y < m {
		ans = append(ans, [2]int{c1x, c1y})
	}

	if c2x, c2y := p2[0]-dx, p2[1]-dy; c2x >= 0 && c2x < n && c2y >= 0 && c2y < m {
		ans = append(ans, [2]int{c2x, c2y})
	}

	return ans
}

func part1(input []string) {
	res := make(map[[2]int]bool)
	points := make(map[rune][][2]int)

	for i, l := range input {
		for j, c := range l {
			if c == '.' {
				continue
			}

			if _, ok := points[c]; !ok {
				points[c] = make([][2]int, 0)
			}

			points[c] = append(points[c], [2]int{i, j})
		}
	}

	n, m := len(input), len(input[0])

	for _, x := range points {
		for i := 0; i < len(x); i++ {
			for j := i + 1; j < len(x); j++ {
				for _, p := range check(x[i], x[j], m, n) {
					res[p] = true
				}
			}
		}
	}

	fmt.Println(len(res))
}

func part2(input []string) {
	res := make(map[[2]int]bool)
	points := make(map[rune][][2]int)

	for i, l := range input {
		for j, c := range l {
			if c == '.' {
				continue
			}
			if _, ok := points[c]; !ok {
				points[c] = make([][2]int, 0)
			}
			points[c] = append(points[c], [2]int{i, j})
		}
	}

	n, m := len(input), len(input[0])

	check2 := func(p1, p2 [2]int) {
		dx, dy := p1[0]-p2[0], p1[1]-p2[1]
		for cx, cy := p1[0], p1[1]; cx >= 0 && cx < n && cy >= 0 && cy < m; cx, cy = cx+dx, cy+dy {
			res[[2]int{cx, cy}] = true
		}

		for cx, cy := p2[0], p2[1]; cx >= 0 && cx < n && cy >= 0 && cy < m; cx, cy = cx-dx, cy-dy {
			res[[2]int{cx, cy}] = true
		}

	}

	for _, x := range points {
		for i := 0; i < len(x); i++ {
			for j := i + 1; j < len(x); j++ {
				check2(x[i], x[j])
			}
		}
	}

	fmt.Println(len(res))
}

func main() {
	lines := util.ReadInput(false)

	part1(lines)
	part2(lines)
}
