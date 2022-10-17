package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

//func solve() {
//	var v0, v1 = make([]int, 0), make([]int, 0)
//	n, q := _readInt(), _readInt()
//	for i := 0; i < n; i++ {
//		x := _readInt()
//		if x%2 == 0 {
//			v0 = append(v0, x)
//		} else {
//			v1 = append(v1, x)
//		}
//	}
//
//	for i := 0; i < q; i++ {
//		var tar, dest []int
//		t, x := _readInt(), _readInt()
//		if t == 0 {
//			tar, dest = v0, v1
//		} else {
//			tar, dest = v1, v0
//		}
//
//		for j := range tar {
//			tar[j] += x
//		}
//	}
//
//}

func solve() {
	n, q := _readInt(), _readInt()
	var s0, n0, s1, n1 int64
	var res int64

	for i := 0; i < n; i++ {
		x := int64(_readInt())
		if x%2 == 0 {
			n0++
			s0 += x
		} else {
			n1++
			s1 += x
		}
		res += x
	}

	for i := 0; i < q; i++ {
		t, x := _readInt(), int64(_readInt())
		if t == 0 {
			res += n0 * x
		} else {
			res += n1 * x
		}
		if x%2 != 0 {
			if t == 0 {
				n1 += n0
				n0 = 0
			} else {
				n0 += n1
				n1 = 0
			}
		}
		_print(res)
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
