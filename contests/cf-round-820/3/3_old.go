package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
//
//func index(r rune) int {
//	return int(r-'a') + 1
//}
//
//func solve() string {
//	var s string
//	fmt.Fscan(in, &s)
//	var ids = make([]int, len(s))
//	var reverse bool
//
//	for i := range s {
//		ids[i] = index([]rune(s)[i])
//	}
//	l, r := ids[0], ids[len(ids)-1]
//
//	if l > r {
//		l, r = r, l
//		reverse = true
//	}
//
//	var res = make([]string, 0, len(ids))
//	var posVals = make([]int, 0, r-l-1)
//	for i := l + 1; i < r; i++ {
//		posVals = append(posVals, i)
//	}
//
//	for _, i := range posVals {
//		for x, j := range ids {
//			if i == j {
//				res = append(res, strconv.Itoa(x+1))
//			}
//		}
//	}
//	if reverse {
//		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
//			res[i], res[j] = res[j], res[i]
//		}
//	}
//
//	resS := strings.Join(res, " ")
//
//	resS = "1 " + resS + strconv.Itoa(len(s)-1)
//	return resS
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
