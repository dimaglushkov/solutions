package main

import "fmt"

// source: https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	alphabet := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	results := make([][]string, len(digits))
	results[0] = make([]string, len(alphabet[rune(digits[0])]))
	for i, r := range alphabet[rune(digits[0])] {
		results[0][i] = string(r)
	}
	if len(digits) == 1 {
		return results[0]
	}

	for i, d := range digits[1:] {
		results[i+1] = make([]string, len(alphabet[d])*len(results[i]))
		counter := 0
		for _, w := range results[i] {
			for _, r := range alphabet[d] {
				results[i+1][counter] = w + string(r)
				counter++
			}
		}
	}

	return results[len(results)-1]
}

func main() {
	// Example 1
	var digits1 string = "23"
	fmt.Println("Expected: [\"ad\",\"ae\",\"af\",\"bd\",\"be\",\"bf\",\"cd\",\"ce\",\"cf\"]	Output: ", letterCombinations(digits1))

	// Example 2
	var digits2 string = ""
	fmt.Println("Expected: []	Output: ", letterCombinations(digits2))

	// Example 3
	var digits3 string = "2"
	fmt.Println("Expected: [\"a\",\"b\",\"c\"]	Output: ", letterCombinations(digits3))

}
