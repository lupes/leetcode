package question_1491_1500

// 1493. 删掉一个元素以后全为 1 的最长子数组
// https://leetcode-cn.com/problems/longest-subarray-of-1s-after-deleting-one-element/
// Topics: 数组

func longestSubarray(nums []int) int {
	left, right, zero, max := 0, 0, 0, 0
	for right < len(nums) {
		if zero < 2 {
			zero += 1 - nums[right]
			if zero < 2 && right-left > max {
				max = right - left
			}
			right++
		} else {
			zero -= 1 - nums[left]
			left++
		}
	}
	return max
}
