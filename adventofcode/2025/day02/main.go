package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func part1(input []string) {
	var res int

	for _, values := range input {
		t := strings.Split(values, "-")
		l, r := t[0], t[1]
		li, ri := util.Atoi(l), util.Atoi(r)

		// round up left to nearest biggest num of even length
		if len(l)%2 != 0 {
			li = int(math.Pow10(len(l)))
			l = strconv.Itoa(li)
		}

		x := util.Atoi(l[:len(l)/2])
		for {
			v := util.Atoi(strconv.Itoa(x) + strconv.Itoa(x))

			if v > ri {
				break
			}

			if v >= li {
				res += v
			}

			x++
		}

	}

	fmt.Println(res)
}

func part2(input []string) {
	var res int

	maxValue := 0

	for _, values := range input {
		t := strings.Split(values, "-")
		l, r := t[0], t[1]
		_, ri := util.Atoi(l), util.Atoi(r)

		maxValue = max(maxValue, ri)
	}

	ids := make(map[int]any)
	maxValueStr := strconv.Itoa(maxValue)

	for i := 1; i <= util.Atoi(maxValueStr[:len(maxValueStr)/2+1]); i++ {
		sb := strings.Builder{}
		s := strconv.Itoa(i)
		sb.WriteString(s)

		for sb.Len() <= len(maxValueStr) {
			sb.WriteString(s)
			ids[util.Atoi(sb.String())] = true
		}
	}

	for _, values := range input {
		t := strings.Split(values, "-")
		l, r := t[0], t[1]
		li, ri := util.Atoi(l), util.Atoi(r)

		for i := li; i <= ri; i++ {
			if _, ok := ids[i]; ok {
				res += i
			}
		}
	}

	fmt.Println(res)
}

//func part2(input []string) {
//	var res int
//
//	for _, values := range input {
//		t := strings.Split(values, "-")
//		l, r := t[0], t[1]
//		li, ri := util.Atoi(l), util.Atoi(r)
//
//		for s := 1; s <= len(r)/2; s++ {
//			for x := util.Atoi(l[:s]); util.Atoi(strconv.Itoa(x)+strconv.Itoa(x)) > ri; x++ {
//				vs := strconv.Itoa(x)
//				for util.Atoi(vs) < li {
//					vs += strconv.Itoa(x)
//				}
//
//				res += util.Atoi(vs)
//				vs += strconv.Itoa(x)
//
//				for util.Atoi(vs) <= ri {
//					res += util.Atoi(vs)
//					vs += strconv.Itoa(x)
//				}
//			}
//		}
//	}
//
//	fmt.Println(res)
//}

func main() {
	lines := util.ReadInput(false)

	data := strings.Split(lines[0], ",")

	part1(data)
	part2(data)
}
