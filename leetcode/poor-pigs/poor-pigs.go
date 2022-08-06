package main

import (
	"fmt"
	"math"
)

// source: https://leetcode.com/problems/poor-pigs/

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}

func main() {
	// Example 1
	var buckets1 int = 1000
	var minutesToDie1 int = 15
	var minutesToTest1 int = 60
	fmt.Println("Expected: 5	Output: ", poorPigs(buckets1, minutesToDie1, minutesToTest1))

	// Example 2
	var buckets2 int = 4
	var minutesToDie2 int = 15
	var minutesToTest2 int = 15
	fmt.Println("Expected: 2	Output: ", poorPigs(buckets2, minutesToDie2, minutesToTest2))

	// Example 3
	var buckets3 int = 4
	var minutesToDie3 int = 15
	var minutesToTest3 int = 30
	fmt.Println("Expected: 2	Output: ", poorPigs(buckets3, minutesToDie3, minutesToTest3))

}
