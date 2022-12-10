package main

import "strconv"

// problem 1

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumValue(strs []string) int {
	res := -1

	for _, s := range strs {
		x, err := strconv.ParseInt(s, 10, 32)
		v := int(x)
		if err != nil {
			v = len(s)
		}
		res = max(res, v)
	}

	return res
}

func main() {
	println()
}
