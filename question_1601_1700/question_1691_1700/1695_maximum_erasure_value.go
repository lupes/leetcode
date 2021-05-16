package question_1691_1700

// 1695. 删除子数组的最大得分
// https://leetcode-cn.com/problems/maximum-erasure-value/
// Topics: 双指针

func maximumUniqueSubarray(nums []int) int {
	var flag, left, right, res, l, tmp = make(map[int]int, len(nums)), 0, 0, 0, len(nums), 0
	for l > right && l > left {
		n := flag[nums[right]]
		if n == 0 {
			tmp, right, flag[nums[right]] = tmp+nums[right], right+1, flag[nums[right]]+1
			if tmp > res {
				res = tmp
			}
		} else {
			tmp, flag[nums[left]], left = tmp-nums[left], flag[nums[left]]-1, left+1
		}
	}
	return res
}
