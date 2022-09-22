package main

func calc(as []int, w int) bool {
	n := len(as)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if dp[i][j] || (j >= as[i] && dp[i+1][j-as[i]]) {
				dp[i+1][j] = true
			}
		}
	}

	return dp[n][w]
}
