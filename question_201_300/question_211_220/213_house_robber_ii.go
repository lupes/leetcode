package question_211_220

// 213. 打家劫舍 II
// https://leetcode-cn.com/problems/house-robber-ii/
// Topics: 动态规划

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := robHelper(nums[:len(nums)-1])
	rightMax := robHelper(nums[1:])
	if max < rightMax {
		max = rightMax
	}
	return max
}

func robHelper(nums []int) int {
	var last, now, tmp = 0, 0, 0
	for _, n := range nums {
		tmp = now
		if last+n > now {
			now = last + n
		}
		last = tmp
	}
	return now
}
