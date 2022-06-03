package main

// source: https://leetcode.com/problems/range-sum-query-2d-immutable/

type NumMatrix [][]int

func Constructor(matrix [][]int) NumMatrix {
	var n, m = len(matrix), len(matrix[0])
	sumMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sumMatrix[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sumMatrix[i][j] = sumMatrix[i-1][j] + sumMatrix[i][j-1] - sumMatrix[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return sumMatrix
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	col1++
	row2++
	col2++
	return (*m)[row2][col2] - (*m)[row2][col1-1] - (*m)[row1-1][col2] + (*m)[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	m := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	m.SumRegion(2, 1, 4, 3) // return 8 (i.e sum of the red rectangle)
	m.SumRegion(1, 1, 2, 2) // return 11 (i.e sum of the green rectangle)
	m.SumRegion(1, 2, 2, 4) // return 12 (i.e sum of the blue rectangle)
}
