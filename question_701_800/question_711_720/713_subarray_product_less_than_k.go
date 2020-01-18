package question_711_720

// 713. 乘积小于K的子数组
// https://leetcode-cn.com/problems/subarray-product-less-than-k
// Topics: 数组 双指针

func numSubarrayProductLessThanK(nums []int, k int) int {
	res, left, tmp := 0, 0, 1
	for right, n := range nums {
		tmp *= n
		for tmp >= k && left <= right {
			tmp /= nums[left]
			left += 1
		}
		res += right - left + 1
	}
	return res
}
