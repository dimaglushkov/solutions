package main

import "fmt"

// source: https://leetcode.com/problems/zigzag-conversion/

// instead of building matrix and inserting zigzagged string
// calculating target indexes from the given string and numRows
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	res := make([]rune, len(s))
	runes := []rune(s)

	var i, d int
	for j := 0; j < len(s); j += 2 * (numRows - 1) {
		res[i] = runes[j]
		i++
	}
	for l := 1; l < numRows-1; l++ {
		res[i] = runes[l]
		i++
		j := l
		for {
			if d = 2 * (numRows - l - 1); j+d >= len(s) {
				break
			}
			j += d
			res[i] = runes[j]
			i++

			if d = 2 * l; j+d >= len(s) {
				break
			}
			j += d
			res[i] = runes[j]
			i++
		}

	}
	for j := numRows - 1; j < len(s); j += 2 * (numRows - 1) {
		res[i] = runes[j]
		i++
	}

	return string(res)
}

func main() {
	// Example 5
	var s5 string = "ABC"
	var numRows5 int = 2
	fmt.Println("Expected: \"ACB\"	Output: ", convert(s5, numRows5))

	var s = "0123456789"
	fmt.Println("Expected: \"0246813579\" 	Output: ", convert(s, 2))
	fmt.Println("Expected: \"0481357926\" 	Output: ", convert(s, 3))
	fmt.Println("Expected: \"0615724839\" 	Output: ", convert(s, 4))

	// Example 2
	var s2 string = "PAYPALISHIRING"
	var numRows2 int = 4
	fmt.Println("Expected: \"PINALSIGYAHRPI\"	Output: ", convert(s2, numRows2))

	// Example 1
	var s1 string = "PAYPALISHIRING"
	var numRows1 int = 3
	fmt.Println("Expected: \"PAHNAPLSIIGYIR\"	Output: ", convert(s1, numRows1))

	// Example 3
	var s3 string = "A"
	var numRows3 int = 1
	fmt.Println("Expected: \"A\"	Output: ", convert(s3, numRows3))

	// Example 4
	var s4 string = "AB"
	var numRows4 int = 1
	fmt.Println("Expected: \"AB\"	Output: ", convert(s4, numRows4))

	// Example 6
	var s6 string = "AB"
	var numRows6 int = 11
	fmt.Println("Expected: \"AB\"	Output: ", convert(s6, numRows6))
}
