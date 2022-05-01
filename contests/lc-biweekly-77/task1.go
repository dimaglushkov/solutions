package main

import "strings"

func countPrefixes(words []string, s string) int {
	var res int
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			res++
		}
	}

	return res
}

func main() {

}
