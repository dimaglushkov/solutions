package day01

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"sort"
	"strconv"
)

func max(x ...int) int {
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] > x[mi] {
			mi = i
		}
	}
	return x[mi]
}

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func DayOne1() {
	cnt, res := 0, 0
	for _, l := range util.ReadInput("day01_1.input") {
		if l == "" {
			res = max(res, cnt)
			cnt = 0
		}
		cnt += atoi(l)
	}
	println(res)
}

func sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}

func DayOne2() {
	cnt := make([]int, 1)
	i := 0
	for _, l := range util.ReadInput("day01_1.input") {
		if l == "" {
			i++
			cnt = append(cnt, 0)
			continue
		}
		cnt[i] += atoi(l)
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	print(sum(cnt[:3]...))
}
