package question_201_210

// 209. 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// Topics: 数组 双指针 二分查找

func minSubArrayLen(s int, nums []int) int {
	nums = append(nums, 0)
	min := len(nums)
	left, right, sum := 0, 0, 0
	for right < len(nums) {
		if sum < s {
			sum += nums[right]
			right++
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= s && min > right-left {
			min = right - left
		}
	}
	if min == len(nums) {
		return 0
	}
	return min
}
