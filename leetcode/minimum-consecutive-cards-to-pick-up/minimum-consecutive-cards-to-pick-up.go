package main

import "fmt"

// source: https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/

func minimumCardPickup(cards []int) int {
	var res = 100001
	var memo = make(map[int]int)
	for i, c := range cards {
		if prevI, ok := memo[c]; ok {
			dist := i - prevI + 1
			if dist < res {
				res = dist
			}
		}
		memo[c] = i
	}
	if res == 100001 {
		return -1
	}
	return res
}

func main() {
	// Example 1
	var cards1 = []int{3, 4, 2, 3, 4, 7}
	fmt.Println("Expected: 4	Output: ", minimumCardPickup(cards1))

	// Example 2
	var cards2 = []int{1, 0, 5, 3}
	fmt.Println("Expected: -1	Output: ", minimumCardPickup(cards2))

}
