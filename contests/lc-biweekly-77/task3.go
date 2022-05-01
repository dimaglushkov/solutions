package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	const (
		guard   = 3
		wall    = 2
		checked = 1
		free    = 0
	)
	var res = 0
	var matrix = make([][]int8, m)
	for i := range matrix {
		matrix[i] = make([]int8, n)
	}

	for _, w := range walls {
		matrix[w[0]][w[1]] = wall
	}

	for _, g := range guards {
		matrix[g[0]][g[1]] = guard
	
	}

	return res
}

func main() {
	countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}})

}
