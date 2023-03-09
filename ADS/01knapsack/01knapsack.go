package knapsack

// Knapsack solve the basic 01knapsack problem:
//
// We are given n items where each item has some weight w and value v associated with it. We are also given a bag
// with capacity c, [i.e., the bag can hold at most c weight in it]. The target is to put the items into the bag
// such that the sum of values associated with them is the maximum possible.
func Knapsack(c int, w, v []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(v)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			if p := j - w[i-1]; p >= 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][p]+v[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][c]
}

// SelectSum
// Simplified 01 knapsack problem: find if it's possible to select any number of elements
// from a given set arr so that sum of the chosen elements is equal to sum
func SelectSum(arr []int, sum int) bool {
	n := len(arr)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= arr[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-arr[i-1]]
			}
		}
	}
	return dp[n][sum]
}
