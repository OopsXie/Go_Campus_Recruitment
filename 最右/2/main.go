/*
小右是一位充满冒险精神的探险家，这次她进入了一座神秘的迷宫。这个迷宫由一个 n×m 的矩阵组成，其中 0 代表小右可以通行的道路，而 1 则是无法穿越的障碍物。小右的任务是从迷宫的入口（左上角 (0, 0)）出发，一路向终点（右下角 (n-1, m-1)）前进。她每次只能向右或向下移动，迷宫充满了挑战和未知。小右需要你的帮助来找到从起点到终点最短路径的长度，如果通往终点的路被阻断，请返回 - 1 来告诉她这条路不存在！
*/

package main

import "fmt"

func shortestPath(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	n, m := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[n-1][m-1] == 1 {
		return -1
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 0
	// 初始化第一列
	for i := 1; i < n; i++ {
		if grid[i][0] == 0 && dp[i-1][0] != -1 {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = -1
		}
	}
	// 初始化第一行
	for j := 1; j < m; j++ {
		if grid[0][j] == 0 && dp[0][j-1] != -1 {
			dp[0][j] = dp[0][j-1] + 1
		} else {
			dp[0][j] = -1
		}
	}
	// 填充dp数组
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if grid[i][j] == 1 {
				dp[i][j] = -1
				continue
			}
			up, left := dp[i-1][j], dp[i][j-1]
			if up == -1 && left == -1 {
				dp[i][j] = -1
			} else if up == -1 {
				dp[i][j] = left + 1
			} else if left == -1 {
				dp[i][j] = up + 1
			} else {
				if up < left {
					dp[i][j] = up + 1
				} else {
					dp[i][j] = left + 1
				}
			}
		}
	}
	return dp[n-1][m-1]
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(shortestPath(grid))
}
