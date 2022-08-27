package main

import "fmt"

// source: https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/

func minimumRecolors(blocks string, k int) int {
	res := 101
	n := len(blocks)
	for i := 0; i <= n-k; i++ {
		cnt := 0
		for j := i; j < i+k; j++ {
			if blocks[j] == 'W' {
				cnt++
			}
		}
		if cnt < res {
			res = cnt
		}
	}

	return res
}

func main() {
	// Example 1
	var blocks1 string = "WBBWWBBWBW"
	var k1 int = 7
	fmt.Println("Expected: 3	Output: ", minimumRecolors(blocks1, k1))

	// Example 2
	var blocks2 string = "WBWBBBW"
	var k2 int = 2
	fmt.Println("Expected: 0	Output: ", minimumRecolors(blocks2, k2))

}
