package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//)
//
//var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
//
//func solve() string {
//	var n, s int
//	fmt.Fscan(in, &n, &s)
//	var tmp int
//	var a = make([]int, 0, n)
//
//	for i := 0; i < n; i++ {
//		fmt.Fscan(in, &tmp)
//		if tmp == 1 {
//			a = append(a, i)
//		}
//	}
//
//	if len(a) < s {
//		return "-1"
//	}
//	if len(a) == s {
//		return "0"
//	}
//
//	var lo, ro = -1, n
//	var res, l, r int
//
//	for d := len(a) - s; d > 0; d-- {
//		l = a[0] - lo
//		r = abs(ro - a[len(a)-1])
//		if l < r {
//			lo = a[0]
//			a = a[1:]
//			res += l
//		} else {
//			ro = a[len(a)-1]
//			a = a[:len(a)-1]
//			res += r
//		}
//	}
//
//	return strconv.FormatInt(int64(res), 10)
//}
//
//func main() {
//	defer out.Flush()
//	var t int
//	for fmt.Fscanln(in, &t); t > 0; t-- {
//		fmt.Fprintln(out, solve())
//	}
//}
//
//// util functions
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
//
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//func makeMatrix(n, m int) [][]int {
//	x := make([][]int, n)
//	for i := range x {
//		x[i] = make([]int, m)
//	}
//	return x
//}
