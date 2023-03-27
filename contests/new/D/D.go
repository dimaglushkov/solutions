package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func _readInt() int {
	var x int
	fmt.Fscan(in, &x)
	return x
}
func _readStr() string {
	var x string
	fmt.Fscan(in, &x)
	return x
}
func _readArr(x ...int) (int, []int) {
	var n int
	if len(x) == 0 {
		fmt.Fscan(in, &n)
	} else {
		n = x[0]
	}
	var v = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	return n, v
}
func _print(x ...any) {
	fmt.Fprintln(out, x...)
}
func _printIntArr(x []int) {
	for _, v := range x {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
func _printStrArr(x []string) {
	for _, v := range x {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
func _yes() {
	fmt.Fprintln(out, "Yes")
}
func _no() {
	fmt.Fprintln(out, "No")
}
func _max(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] > x[mi] {
			mi = i
		}
	}
	return x[mi]
}
func _min(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] < x[mi] {
			mi = i
		}
	}
	return x[mi]
}
func _sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}
func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func _sortStr(x string) string {
	r := []rune(x)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
func _reverseInts(x []int) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}
func _reverseStr(s string) string {
	x := []rune(s)
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
	return string(x)
}
func _makeMatrix(n, m int) [][]int {
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, m)
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		solve()
	}
	out.Flush()
}

// source: _
func solve() {
	n := 0
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	maxSum := 0
	maxV, minV := math.MinInt, math.MaxInt
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] > 0 {
			maxSum += arr[i]
		}
		maxV = max(maxV, arr[i])
		minV = min(minV, arr[i])
	}

	if maxSum < maxV-minV {
		_yes()
		_printIntArr(arr)
		return
	}

	sort.Ints(arr)
	maxSum = 0
	l, r := 0, n-1
	ans := make([]int, n)
	for i := range ans {
		if _abs(maxSum+arr[l]) < _abs(maxSum+arr[r]) {
			maxSum += arr[l]
			ans[i] = arr[l]
			l++
		} else {
			maxSum += arr[r]
			ans[i] = arr[r]
			r--
		}

		if _abs(maxSum) >= maxV-minV {
			_no()
			return
		}
	}
	_yes()
	_printIntArr(ans)

}
