package day01

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"sort"
	"strconv"
)

func DayOne2() {
	cnt := make([]int, 1)
	i := 0
	atoi := func(x string) int {
		y, _ := strconv.ParseInt(x, 10, 32)
		return int(y)
	}
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
	sum := func(x ...int) int {
		res := 0
		for _, i := range x {
			res += i
		}
		return res
	}
	print(sum(cnt[:3]...))
}
