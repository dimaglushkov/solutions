package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
var mem = make(map[int]map[int]bool)

func gcd(a, b int) int {
	if mem[a][b] {
		return 1
	}

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	if a == 1 {
		if _, ok := mem[a]; !ok {
			mem[a] = make(map[int]bool)
		}
		if _, ok := mem[b]; !ok {
			mem[b] = make(map[int]bool)
		}
		mem[a][b] = true
		mem[b][a] = true
	}
	return a
}

func solve() {
	n, v := _readArr()
	max := -3
	for i := n - 1; i >= 0; i-- {
		for j := i; j >= 0 && i+j > max; j-- {
			if (v[j] == 1 || gcd(v[i], v[j]) == 1) && i+j > max {
				max = i + j
			}
		}
	}
	_print(max + 2)
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
