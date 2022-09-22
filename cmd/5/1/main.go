package main

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func calc(as [][3]int) int {
	n := len(as)
	dp := make([][3]int, n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+as[i][k])
			}
		}
	}

	res := 0
	for i := 0; i < 3; i++ {
		chmax(&res, dp[n][i])
	}

	return res
}
