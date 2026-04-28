package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var t int
	fmt.Scan(&t)

	if t == 0 {
		fmt.Println("YES")
	}
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		left := make([]int, n)
		right := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&left[i])
		}
		for i := 0; i < n; i++ {
			fmt.Scan(&right[i])
		}

		if n == 0 {
			fmt.Println("YES")
			continue
		}

		isBalance, _ := calulate(0, left, right, n)

		if isBalance {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func calulate(root int, left []int, right []int, n int) (bool, int) {
	if root == -1 {
		return true, 0
	}

	if root < 0 || root >= n {
		return true, 0
	}

	leftBalance, leftHeight := calulate(left[root], left, right, n)
	if !leftBalance {
		return false, 0
	}

	rightBalance, rightHeight := calulate(right[root], left, right, n)
	if !rightBalance {
		return false, 0
	}

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false, 0
	}

	return true, max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
