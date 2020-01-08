package question_671_680

// 674. 最长连续递增序列
// https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
// Topics: 数组

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var max, start, last = 1, -1, nums[0]
	for i, n := range nums[1:] {
		if n > last {
			if i-start+1 > max {
				max = i - start + 1
			}
			last = n
		} else {
			start, last = i, n
		}
	}
	return max
}
