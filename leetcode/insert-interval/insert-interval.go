package main

import "fmt"

// source: https://leetcode.com/problems/insert-interval/
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	i := 0

	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// Example 1
	var intervals1 = [][]int{{1, 3}, {6, 9}}
	var newInterval1 = []int{2, 5}
	fmt.Println("Expected: [[1,5],[6,9]]	Output: ", insert(intervals1, newInterval1))

	// Example 2
	var intervals2 = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	var newInterval2 = []int{4, 8}
	fmt.Println("Expected: [[1,2],[3,10],[12,16]]	Output: ", insert(intervals2, newInterval2))

}
