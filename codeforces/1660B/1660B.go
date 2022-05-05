package main

import (
	"fmt"
)

// source: https://codeforces.com/contest/1660/problem/B?locale=en
func solve() {
	var n int
	fmt.Scanln(&n)
	var candies = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&candies[i])
	}
	var prevMax = 0
	for {
		var maxId = prevMax
		for i := range candies {
			if candies[i] >= candies[maxId] && i != prevMax {
				maxId = i
			}
		}
		if candies[maxId] == 0 {
			fmt.Println("YES")
			return
		}
		if maxId == prevMax {
			fmt.Println("NO")
			return
		}
		candies[maxId]--
		prevMax = maxId
	}

}

func main() {
	var t, n int
	fmt.Scanln(&t)
	var cases = make([][]int, t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&n)
		cases[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&cases[i][j])
		}
	}

testcase:
	for j := 0; j < t; j++ {
		n = len(cases[j])
		var prevMax = 0
		for {
			var maxId = prevMax
			for i := range cases[j] {
				if cases[j][i] >= cases[j][maxId] && i != prevMax {
					maxId = i
				}
			}
			if cases[j][maxId] == 0 {
				fmt.Println("YES")
				continue testcase
			}
			if maxId == prevMax {
				fmt.Println("NO")
				continue testcase
			}
			cases[j][maxId]--
			prevMax = maxId
		}
	}
}
