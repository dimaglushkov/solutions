package main

import "fmt"

// source: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res = make([]bool, len(candies))
	var max int
	for _, k := range candies {
		if k > max {
			max = k
		}
	}

	for i, k := range candies {
		if k+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

func main() {
	// Example 1 
	var candies1 = []int{2, 3, 5, 1, 3}
	var extraCandies1 int = 3
	fmt.Println("Expected: [true,true,true,false,true] 	Output: ", kidsWithCandies(candies1, extraCandies1))

	// Example 2 
	var candies2 = []int{4, 2, 1, 1, 2}
	var extraCandies2 int = 1
	fmt.Println("Expected: [true,false,false,false,false] 	Output: ", kidsWithCandies(candies2, extraCandies2))

	// Example 3 
	var candies3 = []int{12, 1, 12}
	var extraCandies3 int = 10
	fmt.Println("Expected: [true,false,true]	Output: ", kidsWithCandies(candies3, extraCandies3))

}
