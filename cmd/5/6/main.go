package main

const inf = 1 << 29

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func calc(as []int, ms []int, w int) bool {
	n := len(as)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if dp[i][j] < inf {
				chmin(&dp[i+1][j], 0)
			}
			if j >= as[i] && dp[i+1][j-as[i]] < ms[i] {
				chmin(&dp[i+1][j], dp[i+1][j-as[i]]+1)
			}
		}
	}

	return dp[n][w] < inf
}
