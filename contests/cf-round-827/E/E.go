package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() {
	n, q := _readInt(), _readInt()
	var a = make([]int, n+2)
	for i := 1; i <= n; i++ {
		a[i] = a[i-1] + _readInt()
	}
	a[n+1] = a[n]

	var r = make(map[int]int, q)
	for i := 0; i < q; i++ {
		k := _readInt()
		if x, ok := r[k]; ok {
			fmt.Printf("%d ", a[x])
			continue
		}
		m := 0
		for j := range r {
			if j < k && j > m {
				m = j
			}
		}
		r[k] = r[m]
		for j := r[m]; j < n && a[j+1]-a[j] <= k; j++ {
			r[k]++
		}

		fmt.Printf("%d ", a[r[k]])
	}
	fmt.Println()

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
