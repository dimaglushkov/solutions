package day04

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"strconv"
	"strings"
)

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func DayFour1() {
	res := 0

	for _, l := range util.ReadInput("day04_1.input") {
		s := strings.Split(l, ",")
		x1, x2 := strings.Split(s[0], "-"), strings.Split(s[1], "-")
		a, b, c, d := atoi(x1[0]), atoi(x1[1]), atoi(x2[0]), atoi(x2[1])
		if a <= c && b >= d || a >= c && b <= d {
			res++
		}
	}

	println(res)
}

func DayFour2() {
	res := 0

	for _, l := range util.ReadInput("day04_1.input") {
		s := strings.Split(l, ",")
		x1, x2 := strings.Split(s[0], "-"), strings.Split(s[1], "-")
		s1, e1, s2, e2 := atoi(x1[0]), atoi(x1[1]), atoi(x2[0]), atoi(x2[1])
		if (s1 >= s2 && s1 <= e2) || (e1 >= s2 && e1 <= e2) || (s2 >= s1 && s2 <= e1) || (e2 >= s1 && e2 <= e1) {
			res++
		}
	}

	println(res)
}
