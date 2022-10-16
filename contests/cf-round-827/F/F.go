package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	s, t := strings.Builder{}, strings.Builder{}
	s.WriteString("a")
	t.WriteString("a")
	n := _readInt()
	for i := 0; i < n; i++ {
		id, k, x := _readInt(), _readInt(), _readStr()
		var target *strings.Builder
		if id == 1 {
			target = &s
		} else {
			target = &t
		}
		for j := 0; j < k; j++ {
			target.WriteString(x)
		}

		ss, st := _sortStr(s.String()), _sortStr(t.String())
		if strings.Compare(ss, _reverseStr(st)) == -1 {
			_print("YES")
		} else {
			_print("NO")
		}
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
func _readArr() (int, []int) {
	var n int
	fmt.Fscan(in, &n)
	var v = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
	}
	return n, v
}
func _print(x ...any) {
	fmt.Fprintln(out, x...)
}
func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
