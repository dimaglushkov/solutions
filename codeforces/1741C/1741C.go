package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)

	var sum int
	v := make([]int, n)
	for i := range v {
		fmt.Fscan(in, &v[i])
		sum += v[i]
	}

	var res = n
	for numOfGroups := 1; numOfGroups <= n; numOfGroups++ {
		if sum%numOfGroups != 0 {
			continue
		}

		target := sum / numOfGroups
		cnt, maxCnt, curSum := 0, 0, 0
		for j := 0; j < n; j++ {
			curSum += v[j]
			cnt++
			if curSum > target {
				maxCnt = n
				break
			}
			if curSum == target {
				if cnt > maxCnt {
					maxCnt = cnt
				}
				curSum = 0
				cnt = 0
			}
		}
		if maxCnt < res {
			res = maxCnt
		}
	}

	return strconv.Itoa(res)
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
