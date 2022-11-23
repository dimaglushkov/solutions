package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func solve() string {
	var n, m int
	fmt.Fscan(in, &n, &m)
	var words = make([]string, n)
	for i := range words {
		fmt.Fscan(in, &words[i])
	}
	var res = math.MaxInt
	var temp int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			temp = 0
			for k := 0; k < m; k++ {
				if words[i][k] > words[j][k] {
					temp += int(words[i][k] - words[j][k])
				} else {
					temp += int(words[j][k] - words[i][k])
				}
			}
			if temp < res {
				res = temp
			}
		}
	}

	return strconv.FormatInt(int64(res), 10)
}

func main() {
	defer out.Flush()
	var t int
	for fmt.Fscanln(in, &t); t > 0; t-- {
		fmt.Fprintln(out, solve())
	}
}
