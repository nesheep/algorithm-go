package main

func calc(ps []int) int {
	n := len(ps)
	w := 10000
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, w+1)
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if dp[i][j] || (j >= ps[i] && dp[i][j-ps[i]]) {
				dp[i+1][j] = true
			}
		}
	}

	res := 0
	for j := range dp[n] {
		if dp[n][j] {
			res++
		}
	}

	return res
}
