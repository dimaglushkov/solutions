package main

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
)

func dayEight1() {
	//lines := util.ReadInput("day08_example.input")
	lines := util.ReadInput("day08.input")
	n, m := len(lines), len(lines[0])
	matrix := make([][]int, n)
	vis := make([][]bool, n)

	for i, l := range lines {
		matrix[i] = make([]int, m)
		vis[i] = make([]bool, m)

		for j, b := range l {
			matrix[i][j] = int(b - '0')
		}
	}

	for i := range matrix {
		m1, m2 := -1, -1
		for j := range matrix[i] {
			if x := matrix[i][j]; x > m1 {
				vis[i][j] = true
				m1 = x
			}
			if x := matrix[i][m-j-1]; x > m2 {
				vis[i][m-j-1] = true
				m2 = x
			}
		}
	}

	for j := range matrix[0] {
		m1, m2 := -1, -1
		for i := range matrix {
			if x := matrix[i][j]; x > m1 {
				vis[i][j] = true
				m1 = x
			}
			if x := matrix[n-i-1][j]; x > m2 {
				vis[n-i-1][j] = true
				m2 = x
			}
		}
	}

	res := 0
	for i := range vis {
		for j := range vis[i] {
			if vis[i][j] {
				res++
			}
		}
	}

	println(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dayEight2() {
	//lines := util.ReadInput("day08_example.input")
	lines := util.ReadInput("day08.input")
	n, m := len(lines), len(lines[0])
	matrix := make([][]int, n)
	res := 0

	for i, l := range lines {
		matrix[i] = make([]int, m)
		for j, b := range l {
			matrix[i][j] = int(b - '0')
		}
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			x := matrix[i][j]
			c := [4]int{}

			// up
			p := 0
			for i-c[p]-1 >= 0 && matrix[i-c[p]-1][j] < x {
				c[p]++
			}
			if i-c[p]-1 >= 0 {
				c[p]++
			}

			// down
			p = 1
			for i+c[p]+1 < n && matrix[i+c[p]+1][j] < x {
				c[p]++
			}
			if i+c[p]+1 < n {
				c[p]++
			}

			// left
			p = 2
			for j-c[p]-1 >= 0 && matrix[i][j-c[p]-1] < x {
				c[p]++
			}
			if j-c[p]-1 >= 0 {
				c[p]++
			}

			// right
			p = 3
			for j+c[p]+1 < m && matrix[i][j+c[p]+1] < x {
				c[p]++
			}
			if j+c[p]+1 < m {
				c[p]++
			}

			tmp := 1
			for _, y := range c {
				tmp *= y
			}

			res = max(res, tmp)

		}
	}

	println(res)

}

func main() {
	dayEight1()
	dayEight2()
}
