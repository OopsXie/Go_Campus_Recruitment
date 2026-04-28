package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	results := make([]int, T)
	for j := 0; j < T; j++ {
		var n int
		fmt.Scan(&n)
		h := make([]int, n)

		for i := range h {
			fmt.Scan(&h[i])
		}

		results[j] = maxLong(h)
	}
	for _, k := range results {
		fmt.Println(k)
	}
}

func maxLong(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	dpReverse := make([]int, n)

	for i := range dpReverse {
		dpReverse[i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if arr[i] < arr[j] {
				dpReverse[i] = max(dpReverse[i], dpReverse[j]+1)
			}
		}
	}
	maxlen := 0

	for i := 0; i < n; i++ {
		if i > 0 && i < n-1 {
			if arr[i-1] < arr[i+1] {
				maxlen = max(maxlen, dp[i-1]+dpReverse[i+1]+1)
			}
		}

		if i > 0 {
			maxlen = max(maxlen, dp[i-1]+1)
		}
		if i < n-1 {
			maxlen = max(maxlen, dpReverse[i+1]+1)
		}
	}

	maxlen = max(maxlen, dpReverse[1]+1)
	if n > 1 {
		maxlen = max(maxlen, dp[n-2]+1)
	}

	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
