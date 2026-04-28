/*
一条商业街的 AI 摄像头每分钟会记录一次当前经过的人数，返回为一个非负整数数组 nums，nums [i] 表示第 i 分钟的人数。现在需要统计在任意连续 k 分钟内总经过人数的最大值，以判断是否触发人流预警。
示例：输入: nums = [2, 1, 5, 1, 2, 3, 4, 1, 2, 1], k=3输出: 9
*/

package main

import "fmt"

func maxConsecutiveSum(nums []int, k int) int {
	if len(nums) < k || k <= 0 {
		return 0
	}
	// 计算初始窗口和
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	// 滑动窗口
	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func main() {
	nums := []int{2, 1, 5, 1, 2, 3, 4, 1, 2, 1}
	k := 3
	fmt.Println(maxConsecutiveSum(nums, k))
}
