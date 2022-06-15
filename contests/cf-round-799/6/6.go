package main

import (
	"bufio"
	"fmt"
	"os"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

// 10 + 3
// 9 + 4
// 8 + 5
// 7 + 6
// 3 13 23
// 3: 1 + 1 + 1; 0 + 1 + 2; 0 + 0 + 3;

// 13: 0 + 4 + 9; 0 + 5 + 8; 0 + 6 + 7; 1 + 3 + 9; 2 + 2 + 9; 2 + 3 + 8; 2 + 4 + 7; 2 + 5 + 6;

// 23: 7 + 7 + 9; 6 + 8 + 9; 5 + 9 + 9;

func solve() string {
	var n int
	fmt.Fscan(in, &n)
	var v [10]int
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &tmp)
		if v[tmp%10] < 3 {
			v[tmp%10]++
		}
	}

	var nums = make([]int, 0, 300)
	for i, j := range v {
		for k := 0; k < j; k++ {
			nums = append(nums, i)
		}
	}

	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if (nums[i]+nums[j]+nums[k])%10 == 3 {
					return "YES"
				}
			}
		}
	}
	return "NO"
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}

// util functions
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeMatrix(n, m int) [][]int {
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, m)
	}
	return x
}
