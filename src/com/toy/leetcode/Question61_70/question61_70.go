package Question61_70

func climbStairs(n int) int {

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < n-1; i++ {
		dp[i+1] = dp[i] + dp[i-1]
	}
	return dp[n]
}
