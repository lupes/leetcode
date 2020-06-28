package question_151_160

// 154. 寻找旋转排序数组中的最小值 II
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
// Topics: 数组 二分查找

func findMin2(nums []int) int {
	l, r, c := 0, len(nums)-1, 0
	for r > l {
		c = (l + r) / 2
		if nums[c] > nums[r] {
			l = c + 1
		} else if nums[c] < nums[r] {
			r = c
		} else {
			r--
		}
	}
	return nums[r]
}
