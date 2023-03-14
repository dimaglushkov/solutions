package main

import (
	"bufio"
	"fmt"
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
func _printArr(x []any) {
	for _, v := range x {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprintln(out)
}
func _yes() {
	fmt.Println("YES")
}
func _no() {
	fmt.Println("NO")
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

func main() {
	solve()
	out.Flush()
}

type pair struct {
	x, y int
}

type queue []pair

func (q *queue) push(x, y int) {
	*q = append(*q, pair{x, y})
}
func (q *queue) pop() (int, int) {
	x := (*q)[0]
	*q = (*q)[1:]
	return x.x, x.y
}

func solve() {
	n := _readInt()
	coords := make([]pair, n)
	for i := range coords {
		coords[i] = pair{_readInt(), _readInt()}
	}
	k := _readInt()
	src, dst := _readInt()-1, _readInt()-1

	var q queue
	visited := make([]bool, n)
	q.push(src, 0)

	for len(q) > 0 {
		x, y := q.pop()
		if visited[x] {
			continue
		}
		if x == dst {
			_print(y)
			return
		}
		visited[x] = true
		for i := range coords {
			if !visited[i] && _abs(coords[x].x-coords[i].x)+_abs(coords[x].y-coords[i].y) <= k {
				q.push(i, y+1)
			}
		}
	}
	_print(-1)
}
