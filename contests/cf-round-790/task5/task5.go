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

func solve() string {
	var n, q int
	fmt.Fscan(in, &n, &q)
	var candies = make([]int, n)
	for i := range candies {
		fmt.Fscan(in, &candies[i])
	}

	var queryMemo map[int]int
	var memo = make(map[int]int)
	sort.Ints(candies)
	var util func(cur []int, sum, target, i int)
	util = func(cur []int, sum, target, i int) {
		if _, ok := memo[target]; ok {
			queryMemo[target] = memo[target]
			return
		}

		if sum >= target {
			queryMemo[target] = len(cur)
			return
		}

		for j := i + 1; j < len(candies); j++ {
			cur = append(cur, candies[j])
			util(cur, sum+candies[j], target, j)
			cur = cur[:len(cur)-1]
		}

	}

	var query int
	var res = make([]int, q)

	for i := 0; i < q; i++ {
		queryMemo = make(map[int]int)
		fmt.Fscan(in, &query)
		util([]int{}, 0, query, -1)
		if val, ok := queryMemo[query]; !ok {
			res[i] = -1
		} else {
			res[i] = val
		}
		for k, v := range queryMemo {
			memo[k] = v
		}
	}

	var strs = make([]string, q)
	for i := range res {
		strs[i] = strconv.FormatInt(int64(res[i]), 10)
	}

	return strings.Join(strs, "\n")
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}
