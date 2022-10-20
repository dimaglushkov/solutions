package main

import "fmt"

// source: https://leetcode.com/problems/integer-to-roman/

func intToRoman(num int) string {
	var res string
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	r := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	for num > 0 {
		for v[i] > num {
			i++
		}
		num -= v[i]
		res += r[i]
	}
	return res
}

func main() {
	// Example 1
	var num1 int = 3
	fmt.Println("Expected: III	Output: ", intToRoman(num1))

	// Example 2
	var num2 int = 58
	fmt.Println("Expected: LVIII	Output: ", intToRoman(num2))

	// Example 3
	var num3 int = 1994
	fmt.Println("Expected: MCMXCIV	Output: ", intToRoman(num3))

}
