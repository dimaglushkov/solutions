package main

import "fmt"

// source: https://leetcode.com/problems/minimum-genetic-mutation/

func isSiblings(s, e string) bool {
	if len(s) != len(e) || s == e {
		return false
	}
	mut := 0
	for i := 0; i < 8; i++ {
		if s[i] != e[i] {
			mut++
		}
	}
	return mut == 1
}

func minMutation(start string, end string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}

	var si, ei = -1, -1

	inBank := false
	for i, w := range bank {
		if w == start {
			inBank = true
			si = i
		}
	}
	if !inBank {
		si = len(bank)
		bank = append(bank, start)
	}

	n := len(bank)
	g := make([][]bool, n)

	for i, w := range bank {
		if w == start {
			si = i
		}
		if w == end {
			ei = i
		}

		g[i] = make([]bool, n)
		for j, sw := range bank {
			if isSiblings(w, sw) {
				g[i][j] = true
			}
		}
	}

	if ei == -1 {
		return -1
	}

	distance := make([]int, n)
	for i := range distance {
		distance[i] = -1
	}

	queue := make([]int, 0)
	push := func(x int) {
		queue = append(queue, x)
	}
	pop := func() int {
		x := queue[0]
		queue = queue[1:]
		return x
	}

	distance[si] = 0
	push(si)
	for len(queue) != 0 {
		x := pop()
		for i := range g[x] {
			if g[x][i] && distance[i] == -1 {
				distance[i] = distance[x] + 1
				push(i)
			}
		}

	}

	return distance[ei]
}

func main() {
	var start string = "AACCTTGG"
	var end string = "AATTCCGG"
	var bank = []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}
	fmt.Println("Expected: -1	Output: ", minMutation(start, end, bank))

	// Example 1
	var start1 string = "AACCGGTT"
	var end1 string = "AACCGGTA"
	var bank1 = []string{"AACCGGTA"}
	fmt.Println("Expected: 1	Output: ", minMutation(start1, end1, bank1))

	// Example 2
	var start2 string = "AACCGGTT"
	var end2 string = "AAACGGTA"
	var bank2 = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	fmt.Println("Expected: 2	Output: ", minMutation(start2, end2, bank2))

	// Example 3
	var start3 string = "AAAAACCC"
	var end3 string = "AACCCCCC"
	var bank3 = []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	fmt.Println("Expected: 3	Output: ", minMutation(start3, end3, bank3))

}
