package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/remove-digit-from-number-to-maximize-result/

func removeDigit(number string, digit byte) string {
	var variants = []string{}

	for i := range number {
		if number[i] == digit {
			variants = append(variants, number[:i]+number[i+1:])
		}
	}

	sort.Strings(variants)

	return variants[len(variants)-1]
}

func main() {
	// Example 1
	var number1 string = "123"
	var digit1 byte = '3'
	fmt.Println("Expected: \"12\"	Output: ", removeDigit(number1, digit1))

	// Example 2
	var number2 string = "1231"
	var digit2 byte = '1'
	fmt.Println("Expected: \"231\"	Output: ", removeDigit(number2, digit2))

	// Example 3
	var number3 string = "551"
	var digit3 byte = '5'
	fmt.Println("Expected: \"51\"	Output: ", removeDigit(number3, digit3))

}
