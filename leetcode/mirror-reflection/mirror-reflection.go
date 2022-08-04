package main

import "fmt"

// source: https://leetcode.com/problems/mirror-reflection/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mirrorReflection(p int, q int) int {
	var res = -1
	var up = true
	var x, y int
	for res == -1 {
		switch x {
		case 0:
			x = p
		case p:
			x = 0
		}

		if up {
			y += q
			if y > p {
				y = 2*p - y
				up = false
			}
		} else {
			y -= q
			if y < 0 {
				y = abs(y)
				up = true
			}
		}

		switch {
		case x == p && y == 0:
			res = 0
		case x == p && y == p:
			res = 1
		case x == 0 && y == p:
			res = 2
		}
	}
	return res
}

func main() {
	// Example 0
	var p0 int = 4
	var q0 int = 3
	fmt.Println("Expected: 2	Output: ", mirrorReflection(p0, q0))

	// Example 1
	var p1 int = 2
	var q1 int = 1
	fmt.Println("Expected: 2	Output: ", mirrorReflection(p1, q1))

	// Example 2
	var p2 int = 3
	var q2 int = 1
	fmt.Println("Expected: 1	Output: ", mirrorReflection(p2, q2))

	// Example 3
	var p3 int = 3
	var q3 int = 2
	fmt.Println("Expected: 0	Output: ", mirrorReflection(p3, q3))

}
