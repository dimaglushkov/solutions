package day01

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"strconv"
)

func DayOne1() {
	cnt, res := 0, 0
	atoi := func(x string) int {
		y, _ := strconv.ParseInt(x, 10, 32)
		return int(y)
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for _, l := range util.ReadInput("day01_1.input") {
		if l == "" {
			res = max(res, cnt)
			cnt = 0
		}
		cnt += atoi(l)
	}
	println(res)
}
