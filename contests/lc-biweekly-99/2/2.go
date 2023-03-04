package main

func main() {

}

func coloredCells(n int) int64 {
	dp := make([]int64, n+1)
	dp[1] = 1
	for i := int64(2); i <= int64(n); i++ {
		dp[i] = dp[i-1] + (i-1)*4
	}
	return dp[n]
}
