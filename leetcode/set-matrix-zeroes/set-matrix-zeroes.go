package main

// source: https://leetcode.com/problems/set-matrix-zeroes/
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	zrows, zcols := make([]bool, n), make([]bool, m)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				zrows[i] = true
				zcols[j] = true
			}
		}
	}

	for i := range zrows {
		if zrows[i] {
			for j := 0; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := range zcols {
		if zcols[j] {
			for i := 0; i < n; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}
