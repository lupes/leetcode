package question_491_500

// 494. 目标和
// https://leetcode-cn.com/problems/target-sum/
// Topics: 深度优先搜索 动态规划

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 && S == 0 {
		return 1
	} else if len(nums) == 0 && S != 0 {
		return 0
	}
	return findTargetSumWays(nums[1:], S-nums[0]) + findTargetSumWays(nums[1:], S+nums[0])
}
