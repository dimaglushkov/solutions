package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

// source: https://codeforces.com/contest/1768/problem/C
func solve() {
	n, arr := _readArr()
	v1, v2 := make(map[int]bool, n), make(map[int]bool, n)
	r1, r2 := make([]int, n), make([]int, n)

	for i, v := range arr {
		if !v1[v] {
			v1[v] = true
			r1[i] = v
		} else if !v2[v] {
			v2[v] = true
			r2[i] = v
		} else {
			_print("NO")
			return
		}
	}
loop:
	for i := range arr {
		if r1[i] == 0 {
			for x := arr[i]; x > 0; x-- {
				if !v1[x] {
					r1[i] = x
					v1[x] = true

					continue loop
				}
			}
			_print("NO")
			return
		} else if r2[i] == 0 {
			for x := arr[i]; x > 0; x-- {
				if !v2[x] {
					r2[i] = x
					v2[x] = true
					continue loop
				}
			}
			_print("NO")
			return
		}
	}
	for i := 1; i <= n; i++ {
		if !v1[i] || !v2[i] {
			_print("NO")
			return
		}
	}
	_print("YES")
	for _, v := range r1 {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
	for _, v := range r2 {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
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
