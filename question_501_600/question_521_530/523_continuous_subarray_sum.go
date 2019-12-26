package question_521_530

// 523. 连续的子数组和
// https://leetcode-cn.com/problems/continuous-subarray-sum
// Topics: 数学 动态规划

func checkSubarraySum(nums []int, k int) bool {
	for i, n := range nums {
		sum := n
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if (k == 0 && sum == 0) || (k != 0 && sum%k == 0) {
				return true
			}
		}
	}
	return false
}
