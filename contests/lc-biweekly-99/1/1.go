package main

import (
	"sort"
	"strconv"
)

func main() {
	print(splitNum(10))
	print(splitNum(100))
}

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

func itoa(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func sortStr(x string) string {
	r := []rune(x)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func splitNum(num int) int {
	s := sortStr(itoa(num))
	var s1, s2 string

	for i := 0; i < len(s); i++ {
		s1 += string(s[i])
		if i+1 < len(s) {
			i++
			s2 += string(s[i])
		}
	}
	return atoi(s1) + atoi(s2)
}
