package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	B := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&B[i])
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)

		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			dp[i][j] = 1 + dp[i+1][j]

			if B[i] == B[j] {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1])

			} else {
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}

	fmt.Println(dp[0][n-1])
}
