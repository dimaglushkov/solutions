package main

import (
	"fmt"
)

// source: https://codeforces.com/contest/1669/problem/A?locale=en

func main() {
	var t int
	fmt.Scanln(&t)

	var nums = make([]int, t)
	for i := range nums {
		fmt.Scanln(&nums[i])
	}

	for i := range nums {
		switch {
		case nums[i] <= 1399:
			fmt.Println("Division 4")
		case 1400 <= nums[i] && nums[i] <= 1599:
			fmt.Println("Division 3")
		case 1600 <= nums[i] && nums[i] <= 1899:
			fmt.Println("Division 2")
		case 1900 <= nums[i]:
			fmt.Println("Division 1")
		}
	}

}
