package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/baseball-game/

func calPoints1(ops []string) int {
	results := []int{}
	res := 0
	v := 0
	for _, op := range ops {
		num, err := strconv.ParseInt(op, 10, 32)
		switch {
		case err == nil:
			v = int(num)
			res += v
			results = append(results, v)
		case op == "C":
			v = len(results) - 1
			res -= results[v]
			results = results[:v]
		case op == "D":
			v = results[len(results)-1] * 2
			res += v
			results = append(results, v)
		case op == "+":
			v = results[len(results)-1] + results[len(results)-2]
			res += v
			results = append(results, v)
		}

	}
	return res
}

func calPoints(ops []string) int {
	results := []int{}
	for _, op := range ops {
		switch op {
		case "C":
			results = results[:len(results)-1]
		case "D":
			results = append(results, results[len(results)-1]*2)
		case "+":
			results = append(results, results[len(results)-1]+results[len(results)-2])
		default:
			num, _ := strconv.ParseInt(op, 10, 32)
			results = append(results, int(num))
		}
	}
	res := 0
	for _, v := range results {
		res += v
	}
	return res
}

func main() {
	// Example 1
	var ops1 = []string{"5", "2", "C", "D", "+"}
	fmt.Println("Expected: 30	Output: ", calPoints(ops1))

	// Example 2
	var ops2 = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println("Expected: 27	Output: ", calPoints(ops2))

	// Example 3
	var ops3 = []string{"1"}
	fmt.Println("Expected: 1	Output: ", calPoints(ops3))

}
