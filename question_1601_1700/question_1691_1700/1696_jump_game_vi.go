package question_1691_1700

// 1696. 跳跃游戏 VI
// https://leetcode-cn.com/problems/jump-game-vi/
// Topics: 动态规划

func maxResult(nums []int, k int) int {
	var dp, flag, max = make([]int, len(nums)), make(map[int]int), nums[0]
	dp[0] = nums[0]
	flag[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		if i-k-1 >= 0 {
			flag[dp[i-k-1]]--
			if flag[dp[i-k-1]] == 0 {
				delete(flag, dp[i-k-1])
			}
			if dp[i-k-1] == max {
				max = dp[i-k]
				for k := range flag {
					if k > max {
						max = k
					}
				}
			}
		}

		dp[i] = max + nums[i]
		flag[max+nums[i]]++
		if max+nums[i] > max {
			max = max + nums[i]
		}
	}
	return dp[len(nums)-1]
}
