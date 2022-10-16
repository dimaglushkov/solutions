package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/add-binary/

func addBinary(a string, b string) string {
	var na, nb int
	if na, nb = len(a), len(b); na < nb {
		a, b = b, a
		na, nb = nb, na
	}

	var res = make([]rune, na+1)
	var p = na
	var overflow bool
	var next int
	for i := range a {
		next = 0
		if a[na-1-i] == '1' {
			next++
		}
		if nb-1-i >= 0 && b[nb-1-i] == '1' {
			next++
		}
		if overflow {
			next++
		}

		switch next {
		case 0:
			res[p] = '0'
			overflow = false
		case 1:
			res[p] = '1'
			overflow = false
		case 2:
			res[p] = '0'
			overflow = true
		case 3:
			res[p] = '1'
			overflow = true
		}
		p--
	}
	if overflow {
		res[p] = '1'
	} else {
		res = res[1:]
	}
	return string(res)
}

func main() {
	var a string = "0"
	var b string = "0"
	fmt.Println("Expected: \"100\"	Output: ", addBinary(a, b))

	// Example 1
	var a1 string = "11"
	var b1 string = "1"
	fmt.Println("Expected: \"100\"	Output: ", addBinary(a1, b1))

	// Example 2
	var a2 string = "1010"
	var b2 string = "1011"
	fmt.Println("Expected: \"10101\"	Output: ", addBinary(a2, b2))

}
