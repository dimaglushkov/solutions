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
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		solve()
	}
	out.Flush()
}

func check(x []int, size int) int {
	n := len(x)
	i := 0
	cnt := 0
	changed := false
	for i < n {
		cnt++
		if i+x[i]+1 > n && !changed {
			changed = true
			i++
		} else {
			i += x[i] + 1
		}
	}
	if i == n && cnt == size && !changed {
		return 0
	}
	if i == n && (cnt >= size && changed || !changed) {
		return 1
	}
	return 2
}

// source: _
func solve() {
	n, arr := _readArr()
	ans := make([]int, n-1)

	for i := n - 2; i >= 0; i-- {
		ans[i] = check(arr[i+1:], arr[i])
	}

	_printIntArr(ans)
}
