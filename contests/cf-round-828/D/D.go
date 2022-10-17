package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n := _readInt()
	var cnt, res int
	for i := 0; i < n; i++ {
		x := _readInt()
		for x > 0 && x%2 == 0 {
			x /= 2
			cnt++
		}
	}

	if cnt >= n {
		_print(0)
		return
	}

	counts := make([]int, 0)
	for i := 1; i <= n; i++ {
		x := i
		if x%2 != 0 {
			continue
		}
		var tmp int
		for x > 0 && x%2 == 0 {
			x /= 2
			tmp++
		}
		counts = append(counts, tmp)
	}
	sort.Ints(counts)
	_reverseInts(counts)

	for _, v := range counts {
		res++
		cnt += v
		if cnt >= n {
			_print(res)
			return
		}
	}

	_print(-1)
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
