package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func index(r rune) int {
	return int(r-'a') + 1
}

type pair struct {
	r rune
	i int
}

func solve() string {
	var s string
	fmt.Fscan(in, &s)
	var runes = []rune(s)
	var l, r = runes[0], runes[len(runes)-1]
	var reverse bool

	if l > r {
		l, r = r, l
		reverse = true
	}

	var res = make([]pair, 0, len(runes))

	for i, x := range runes {
		if x >= l && x <= r {
			res = append(res, pair{r: x, i: i + 1})
		}
	}

	if reverse {
		sort.Slice(res, func(i, j int) bool {
			if res[i].r < res[j].r {
				return false
			}
			return true
		})
	} else {
		sort.Slice(res, func(i, j int) bool {
			if res[i].r < res[j].r {
				return true
			}
			return false
		})
	}

	str1 := fmt.Sprintf("%d %d\n", abs(int(l-r)), len(res))
	var strArr = make([]string, 1, len(res))
	strArr[0] = "1"
	for i := 0; i < len(res); i++ {
		if res[i].i != 1 && res[i].i != len(runes) {
			strArr = append(strArr, strconv.Itoa(res[i].i))
		}
	}
	strArr = append(strArr, strconv.Itoa(len(runes)))

	return str1 + strings.Join(strArr, " ")
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}

// util functions
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func makeMatrix(n, m int) [][]int {
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, m)
	}
	return x
}
