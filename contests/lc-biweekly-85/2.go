package main

import "strings"

func secondsToRemoveOccurrences(s string) int {
	var cnt int

	for {
		if !strings.Contains(s, "01") {
			break
		}
		cnt++
		s = strings.ReplaceAll(s, "01", "10")
	}

	return cnt
}

func main() {
	println(secondsToRemoveOccurrences("0110101"))
	println(secondsToRemoveOccurrences("001011"))
	println(secondsToRemoveOccurrences("11100"))
	println(secondsToRemoveOccurrences("111"))
	println(secondsToRemoveOccurrences("1"))
	println(secondsToRemoveOccurrences("01"))
	println(secondsToRemoveOccurrences("10"))
}
