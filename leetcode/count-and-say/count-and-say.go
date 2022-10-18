package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://leetcode.com/problems/count-and-say/

func countAndSay(n int) string {
	s := "1"

	for i := 1; i < n; i++ {
		cnt := int64(1)
		p := s[0]
		ns := strings.Builder{}

		for j := 1; j < len(s); j++ {
			if s[j] != p {
				ns.WriteString(strconv.FormatInt(cnt, 10))
				ns.WriteString(string(p))
				cnt = 1
				p = s[j]
			} else {
				cnt++
			}
		}
		ns.WriteString(strconv.FormatInt(cnt, 10))
		ns.WriteString(string(p))
		s = ns.String()
	}

	return s
}

func main() {
	// Example 1
	var n1 int = 1
	fmt.Println("Expected: 1	Output: ", countAndSay(n1))

	// Example 2
	var n2 int = 4
	fmt.Println("Expected: 1211	Output: ", countAndSay(n2))

}
