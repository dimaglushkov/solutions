package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

// source: https://codeforces.com/contest/1760/problem/G
func solve() {
	n, a, b := _readInt(), _readInt(), _readInt()
	graph := make([]map[int]int, n+1)
	for i := 0; i < n-1; i++ {
		u, v, w := _readInt(), _readInt(), _readInt()
		if graph[u] == nil {
			graph[u] = make(map[int]int)
		}
		if graph[v] == nil {
			graph[v] = make(map[int]int)
		}
		graph[u][v] = w
		graph[v][u] = w
	}

	vals := make(map[int]bool)
	var dfs1 func(v, prev, x int)
	dfs1 = func(v, prev, x int) {
		if v == b {
			return
		}
		vals[x] = true
		for u, w := range graph[v] {
			if u == prev {
				continue
			}
			dfs1(u, v, w^x)
		}
	}
	dfs1(a, -1, 0)

	var dfs2 func(v, prev, x int) bool
	dfs2 = func(v, prev, x int) bool {
		if v != b && vals[x] {
			return true
		}
		for u, w := range graph[v] {
			if u == prev {
				continue
			}
			if dfs2(u, v, w^x) {
				return true
			}
		}
		return false
	}

	if dfs2(b, -1, 0) {
		_print("YES")
	} else {
		_print("NO")
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
