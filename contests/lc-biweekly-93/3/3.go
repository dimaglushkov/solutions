package main

// problem 3
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(x ...int) int {
	res := 0
	for _, i := range x {
		res += i
	}
	return res
}

func maxJump(s []int) int {
	n := len(s)

	res := s[1] - s[0]
	for i := 2; i < n; i++ {
		res = max(res, s[i]-s[i-2])
	}
	return res
}

/*func maxJump(stones []int) int {
	res := math.MaxInt

	n := len(stones)
	visited := make([]bool, n)

	var backtrack func(ci, mj, ti int)
	backtrack = func(ci, mj, ti int) {
		if mj >= res {
			return
		}
		if ci == ti && ti == n-1 {
			tr := 0
			prev := n - 1
			for i := n - 2; i >= 0; i-- {
				if !visited[i] || i == 0 {
					tr = max(tr, abs(stones[i]-stones[prev]))
					prev = i
				}
			}
			tr = max(tr, mj)
			res = min(res, tr)
			return
		}
		for i := ci + 1; i < n; i++ {
			visited[i] = true
			backtrack(i, max(mj, abs(stones[ci]-stones[i])), ti)
			visited[i] = false
		}

	}
	backtrack(0, 0, n-1)
	return res
}*/

func main() {
	println(maxJump([]int{0, 20, 116, 117, 280, 283, 289, 327, 384, 392, 489, 497, 564, 597, 648, 722, 737, 906, 918, 976, 1050, 1068, 1128, 1215, 1246, 1306, 1324, 1327, 1336, 1359, 1362, 1369, 1380, 1417, 1575, 1639, 1692, 1759, 1762, 1803, 1883, 1971, 1977, 1993, 2039, 2085, 2088, 2109, 2191, 2440}), 0)
	println(maxJump([]int{0, 2, 5, 6, 7}), 5)
	println(maxJump([]int{0, 2, 5, 6, 7}), 5)
	println(maxJump([]int{0, 3, 9}), 9)
}
