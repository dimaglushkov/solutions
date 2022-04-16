package main

import "fmt"

// source: https://leetcode.com/problems/spiral-matrix-ii/

func generateMatrix(n int) [][]int {
	var rl, rr = 0, n - 1
	var cl, cr = 0, n - 1
	var val = 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for val <= n*n {
		for i := cl; i <= cr; i++ {
			res[rl][i] = val
			val++
		}
		for i := rl + 1; i <= rr; i++ {
			res[i][cr] = val
			val++
		}

		for i := cr - 1; i > cl; i-- {
			res[rr][i] = val
			val++
		}
		for i := rr; i > rl; i-- {
			res[i][cl] = val
			val++
		}

		rl++
		rr--
		cl++
		cr--

	}
	return res
}

func main() {
	// Example 3
	var n3 int = 4
	fmt.Println("Expected: [[1,2,3],[8,9,4],[7,6,5]]	Output: ", generateMatrix(n3))

	// Example 1
	var n1 int = 3
	fmt.Println("Expected: [[1,2,3],[8,9,4],[7,6,5]]	Output: ", generateMatrix(n1))

	// Example 2
	var n2 int = 1
	fmt.Println("Expected: [[1]]	Output: ", generateMatrix(n2))

}
