package main

func main() {
	println(maximumCostSubstring("kqqqqqkkkq", "kq", []int{-6, 6})) //30
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximumCostSubstring(s string, chars string, vals []int) int {
	cost := make(map[byte]int)
	for i := byte('a'); i <= 'z'; i++ {
		cost[i] = int(i-'a') + 1
	}
	for i := range chars {
		cost[chars[i]] = vals[i]
	}

	cur := 0
	ans := 0
	for i := range s {
		cur = max(cur+cost[s[i]], 0)
		ans = max(ans, cur)
	}
	return ans
}
