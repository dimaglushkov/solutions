package knapsack

// Knapsack
// Simplified 01 knapsack problem: find if it's possible to select any number of elements
// from a given set arr so that sum of the chosen elements is equal to sum
func Knapsack(arr []int, sum int) bool {
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
