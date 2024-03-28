package main

// source: https://leetcode.com/problems/task-scheduler/
func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	maxFreq := 0

	for i := range tasks {
		m[tasks[i]]++
		maxFreq = max(m[tasks[i]], maxFreq)
	}

	ans := (maxFreq - 1) * (n + 1)

	for _, v := range m {
		if v == maxFreq {
			ans++
		}
	}

	return max(ans, len(tasks))
}
