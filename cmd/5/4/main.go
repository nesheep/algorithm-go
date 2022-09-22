package main

const inf = 2 << 30

func chmmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func calc(as []int, w int, k int) bool {
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
			chmmin(&dp[i+1][j], dp[i][j])
			if j >= as[i] {
				chmmin(&dp[i+1][j], dp[i][j-as[i]]+1)
			}
		}
	}

	return dp[n][w] <= k
}
