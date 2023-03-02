package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

// source: https://codeforces.com/contest/1800/problem/E1
func solve() {
	n, k := _readInt(), _readInt()
	s, t := _readStr(), _readStr()

	cnt := make(map[rune]int)
	for _, r := range s {
		cnt[r]++
	}
	for _, r := range t {
		cnt[r]--
		if cnt[r] == 0 {
			delete(cnt, r)
		}
	}

	if len(cnt) != 0 {
		_print("no")
		return
	}

	if n >= 2*k {
		_print("yes")
		return
	}

	l, r := 0, n-1
	for i := 0; i < n-k; i++ {
		l++
		r--
	}
	if s[l:r+1] == t[l:r+1] {
		_print("yes")
	} else {
		_print("no")
	}
}

//func solve() {
//	n, k := _readInt(), _readInt()
//	s, t := []byte(_readStr()), []byte(_readStr())
//
//	var check func(s1, s2 []byte, i int) bool
//	check = func(s1, s2 []byte, i int) bool {
//		if i == n {
//			return true
//		}
//		if s1[i] == s2[i] {
//			return check(s1, s2, i+1)
//		}
//		for d := 1; i+d*3 < n; i++ {
//			for _, dk := range []int{k, k + 1} {
//				if di := i + dk*d; di < n && s[di] == t[i] {
//					ns1 := make([]byte, n)
//					copy(ns1, s)
//					ns1[i], ns1[di] = ns1[di], ns1[i]
//					if check(ns1, s2, i+1) {
//						return true
//					}
//				}
//			}
//		}
//		return false
//	}
//
//	if check(s, t, 0) {
//		_print("YES")
//	} else {
//		_print("NO")
//	}
//
//}

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
