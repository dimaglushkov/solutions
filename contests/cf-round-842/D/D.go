package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

// Was not able to send in time
// source: https://codeforces.com/contest/1768/problem/D
func solve() {
	n, arr := _readArr()
	cnt := 0
	pos := make(map[int]int, n)
	for i, v := range arr {
		pos[v] = i
	}

	inv := false

	swap := func(i, j int) {
		pos[arr[i]], pos[arr[j]], arr[i], arr[j] = pos[arr[j]], pos[arr[i]], arr[j], arr[i]
		cnt++
	}

	p1, p2 := -1, -1

	for i := 0; i < n-1; i++ {
		if arr[i] == i+2 && arr[i+1] == arr[i]-1 {
			inv = true
			p1, p2 = i, i+1
			break
		}
	}

	for i := 0; i < n; i++ {
		if i == p1 || i == p2 {
			continue
		}
		d := i + 1
		if !inv && i < n-1 && arr[i] == d+1 && arr[i+1] == arr[i]-1 {
			i++
			inv = true
			continue
		}

		if invVal := arr[i] - 1; !inv && arr[i] == d+1 && i < n-1 && pos[invVal] == arr[d]-1 {
			swap(i+1, pos[invVal])
			inv = true
			i++
		} else if arr[i] != d {
			swap(i, pos[d])
		}
	}
	if inv {
		_print(cnt)
	} else {
		_print(cnt + 1)
	}
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		solve()
	}
}

// ====
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
