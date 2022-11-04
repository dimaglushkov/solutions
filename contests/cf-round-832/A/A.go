package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

//void find_min_diff( uint32_t sum_heap1 = 0, uint32_t sum_heap2 = 0, uint8_t i = 0 )
//{
//if (i != size)
//{
//find_min_diff(sum_heap1 + rocks[i], sum_heap2, i + 1);
//find_min_diff(sum_heap1, sum_heap2 + rocks[i], i + 1);
//}
//else
//min_diff = min_diff < sum_heap1 - sum_heap2 ? min_diff : sum_heap1 - sum_heap2;
//}

func solve() {
	_, arr := _readArr()
	s1, s2 := 0, 0
	sort.Ints(arr)
	for _, x := range arr {
		if _abs(s1+x)-s2 > s1-_abs(s2+x) {
			s1 += x
		} else {
			s2 += x
		}
	}

	_print(_abs(s1 - s2))
}

//func solve() {
//	n, arr := _readArr()
//	res := math.MinInt
//	var util func(s1, s2, i int)
//	util = func(s1, s2, i int) {
//		if i != n {
//			util(s1+arr[i], s2, i+1)
//			util(s1, s2+arr[i], i+1)
//		} else {
//			if x := _abs(s1) - _abs(s2); res < x {
//				res = x
//			}
//		}
//	}
//	util(0, 0, 0)
//	_print(res)
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
