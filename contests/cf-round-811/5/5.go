package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	var fives int
	for i := range arr {
		fmt.Fscan(in, &arr[i])
		arr[i] += arr[i] % 10
		if arr[i]%10 == 5 || arr[i]%10 == 0 {
			fives++
		}
	}

	if fives > 0 && fives != n {
		return "NO"
	}

	sort.Ints(arr)
	min := arr[0]
	max := arr[n-1]
	if min == max {
		return "YES"
	}

	if fives == n {
		if max-min == 5 && min%10 == 5 {
			return "YES"
		} else {
			return "NO"
		}
	}

	//dedup
	d := make([]int, 0, n)
	for i := 0; i < n-1; i++ {
		if arr[i] != arr[i+1] {
			d = append(d, arr[i])
		}
	}
	d = append(d, arr[n-1])

	// AHAHHA THAT ACTUALLY WORKED LMAO
	for i := 0; i < n-1; i++ {
		for arr[i]+100000000 < max {
			arr[i] += 100000000
		}
		for arr[i]+10000000 < max {
			arr[i] += 10000000
		}
		for arr[i]+1000000 < max {
			arr[i] += 1000000
		}
		for arr[i]+100000 < max {
			arr[i] += 100000
		}
		for arr[i]+10000 < max {
			arr[i] += 10000
		}
		for arr[i]-max < 10 {
			arr[i] += arr[i] % 10
		}
		if abs(max-arr[i]) == 10 {
			return "NO"
		}
	}

	return "YES"
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
