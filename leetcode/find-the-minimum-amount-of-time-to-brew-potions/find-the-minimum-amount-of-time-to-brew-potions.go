package main

// source: https://leetcode.com/problems/find-the-minimum-amount-of-time-to-brew-potions/
func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	times := make([]int64, n)

	for j := 0; j < m; j++ {
		var curTime int64 = 0
		for i := 0; i < n; i++ {
			if curTime < times[i] {
				curTime = times[i]
			}
			curTime += int64(skill[i]) * int64(mana[j])
		}
		times[n-1] = curTime
		for i := n - 2; i >= 0; i-- {
			times[i] = times[i+1] - int64(skill[i+1])*int64(mana[j])
		}
	}
	return times[n-1]
}
