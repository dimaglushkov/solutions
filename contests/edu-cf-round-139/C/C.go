package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

// source: https://codeforces.com/contest/1766/problem/C
func solve() {
	n := _readInt()
	c := []string{_readStr(), _readStr()}
	type pos struct {
		i, j int
	}

	queue := make([]pos, 0)
	push := func(i, j int) {
		queue = append(queue, pos{i, j})
	}
	pop := func() pos {
		x := queue[0]
		queue = queue[1:]
		return x
	}
	need := strings.Count(c[0], "B") + strings.Count(c[1], "B")

	for _, si := range []int{0, 1} {
		sj := 0
		if c[si][sj] != 'B' {
			continue
		}
		visited := [][]bool{make([]bool, n), make([]bool, n)}
		cnt := 0
		push(si, sj)

		for len(queue) > 0 {
			p := pop()
			visited[p.i][p.j] = true
			cnt++
			if p.i == 0 && c[1][p.j] == 'B' && !visited[1][p.j] {
				push(1, p.j)
			} else if p.i == 1 && c[0][p.j] == 'B' && !visited[0][p.j] {
				push(0, p.j)
			} else if p.j+1 < n && c[p.i][p.j+1] == 'B' && !visited[p.i][p.j+1] {
				push(p.i, p.j+1)
			}
		}
		if need == cnt {
			_print("YES")
			return
		}
	}

	_print("NO")
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
