package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	// if len(firstLine) < 2 {
	// 	panic("invalid first line")
	// }
	N, _ := strconv.Atoi(firstLine[0])
	M, _ := strconv.Atoi(firstLine[1])

	items := make([][3]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		items[i] = [3]int{x, y, z}
	}

	maxScore := 0

	signs := [][]int{{1, 1, 1}, {1, 1, -1}, {1, -1, 1}, {1, -1, -1}, {-1, 1, 1}, {-1, 1, -1}, {-1, -1, 1}, {-1, -1, -1}}

	for _, s := range signs {
		sx, sy, sz := s[0], s[1], s[2]
		dp := make([]int, M+1)
		for i := 1; i <= M; i++ {
			dp[i] = math.MinInt64

		}
		dp[0] = 0

		for _, item := range items {
			x, y, z := item[0], item[1], item[2]
			val := sx*x + sy*y + sz*z
			for k := M; k >= 1; k-- {
				if dp[k-1] != math.MaxInt64 {
					if dp[k] < dp[k-1]+val {
						dp[k] = dp[k-1] + val
					}
				}
			}
		}
		if dp[M] > maxScore {
			maxScore = dp[M]
		}
	}

	fmt.Println(maxScore)

}
